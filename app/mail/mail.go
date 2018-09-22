package main

import (
	"fmt"
	"math/rand"
	"time"
	"os/signal"
	"conf"
)

type MailSrv struct {

}

type MailConfig struct {
	MailEnable string
	Server     string //邮件服务器地址
	SrvPort    string //邮件服务器端口
	User       string //用户名
	Passwd     string //密码
	Proto      string //发送邮件的协议，目前只支持SMTP
	Receiver   string //接收邮件的帐号
	SSL        string
}


func getuserconfig()(string,string,string,string,string) {  //从数据库拿到用户数据
	return "", "", "", "", ""
}

func (this *Mailsever) Securitycode() string {   //发送验证码
	_msg := GetRandomString(6)
	_sbuj := "EasyBlog Security Code"
	_timeout := 10
	_hsot , _port , _user, _pw, _to := geruserconfig() 

	err, _ := SMTPSendStr(_host, _port,_user, _pw, _to,_subj, _msg , _timeout , true)
	if err != nil {
		fmt.Println(err)
	}
	return _msg
}

func GetRandomString(l int) []byte {
	str := "123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return result
}

func init() {
	go func() {
		err := config.MailSrvNotify.Services(new MailSrv)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	osexit := make(chan Os.Signal, 1)
	Signal.Notify(osexit, os.Kill)	
	osexit<-
}