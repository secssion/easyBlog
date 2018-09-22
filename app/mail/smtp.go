package config

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"
	"time"
)

/*
auth
support smtp unencrypted pw
*/
type plainAuth struct {
	identity, username, password string
	host                         string
}

func PlainAuth(identity, username, password, host string) smtp.Auth {
	return &plainAuth{identity, username, password, host}
}
func (a *plainAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	if server.Name != a.host {
		return "", nil, errors.New("wrong host name")
	}
	resp := []byte(a.identity + "\x00" + a.username + "\x00" + a.password)
	return "PLAIN", resp, nil
}
func (a *plainAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		return nil, errors.New("unexpected server challenge")
	}
	return nil, nil
}

type loginAuth struct {
	username, password string
	host               string
}

func LoginAuth(username, password, host string) smtp.Auth {
	return &loginAuth{username, password, host}
}
func (a loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	if server.Name != a.host {
		return "", nil, errors.New("wrong host name")
	}
	return "LOGIN", nil, nil
}
func (a loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		if bytes.EqualFold([]byte("username:"), fromServer) {
			return []byte(a.username), nil
		} else if bytes.EqualFold([]byte("password:"), fromServer) {
			return []byte(a.password), nil
		}
	}
	return nil, nil
}

/*
send str by smtp
support conn/auth timeout & unenc pw auth & ssl trans
*/
func SMTPSendStr(_host, _port,
	_user, _pw, _to,
	_subj, _msg string,
	_timeout int,
	_ssl bool) (error, int) {
	//
	from := mail.Address{"", _user}
	to := mail.Address{"", _to}
	hdr := make(map[string]string)
	hdr["From"] = from.String()
	hdr["To"] = to.String()
	hdr["Subject"] = "=?UTF8?B?" + base64.StdEncoding.EncodeToString([]byte(_subj)) + "?="
	msg := ""
	for k, v := range hdr {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "Content-Type: text/plain; charset=UTF-8\r\n"
	msg += "\r\n" + _msg
	//
	auth := LoginAuth(
		_user,
		_pw,
		_host,
	)
	//connect to the remote server
	var err error
	var clt *smtp.Client
	var sslConn *tls.Conn
	var connChan chan int
	var ret int

	connChan = make(chan int)
	ret = 0

	go func() {
		if _ssl {
			tlsconfig := &tls.Config{
				InsecureSkipVerify: true,
				ServerName:         _host,
			}
			sslConn, err = tls.Dial("tcp", _host+":"+_port, tlsconfig)
			if err == nil {
				clt, err = smtp.NewClient(sslConn, _host)
			}
		} else {
			clt, err = smtp.Dial(_host + ":" + _port)
		}
		connChan <- 1
		if ret != 0 {
			if clt != nil {
				clt.Close()
			}
			if _ssl && sslConn != nil {
				sslConn.Close()
			}
			if connChan != nil {
				close(connChan)
			}
		}
	}()
	if _timeout <= 0 {
		<-connChan
	} else {
		select {
		case <-time.After(time.Second * time.Duration(_timeout)): //连接超时
			ret = 1
			goto END
		case <-connChan:
		}
	}
	if nil == err {
		//login
		go func() {
			err = clt.Auth(auth)
			connChan <- 1
			if ret != 0 {
				clt.Close()
				if _ssl && sslConn != nil {
					sslConn.Close()
				}
				if connChan != nil {
					close(connChan)
				}
			}
		}()
		if _timeout <= 0 {
			<-connChan
		} else {
			select {
			case <-time.After(time.Second * time.Duration(_timeout)): //认证超时
				ret = 1
				goto END
			case <-connChan:
			}
		}
		defer func() {
			clt.Close()
			if _ssl && sslConn != nil {
				sslConn.Close()
			}
			if connChan != nil {
				close(connChan)
			}
		}()
		if err == nil {
			//set the sender
			if err = clt.Mail(_user); err == nil {
				//set the receiver
				if err = clt.Rcpt(_to); err == nil {
					//send the email body
					if writer, err := clt.Data(); err == nil {
						if _, err = writer.Write([]byte(msg)); err == nil {
							if err = writer.Close(); err == nil {
								//send the QUIT command
								err = clt.Quit()
								goto END
							}
						}
					}
					ret = 6
				} else {
					ret = 5
				}
			} else {
				ret = 4
			}
		} else {
			//登陆失败，不支持此服务器
			ret = 3
			if strings.Contains(fmt.Sprintf("%s", err), "authentication failed") { //登陆失败，用户或密码有误
				ret = 31
			}
		}
	} else {
		//连接出错
		ret = 2
		if strings.Contains(fmt.Sprintf("%s", err), "connection refused") { //服务器拒绝连接
			ret = 21
		}
	}

END:
	return err, ret
}
