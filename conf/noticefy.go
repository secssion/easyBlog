package conf

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type NotifyMsgSt struct {
	Cmd  string              //消息类型
	Args interface{}		 //消息内容
}

type NotifyObjectSt struct {
	addr         string
	port         string
	serverMethod string
}

//需要通知的程序注册结构
const (
	NOTIFY_ADDR_SERV = "127.0.0.1"
	NOTIFY_ADDR_NET     = "127.0.0.1"
	NOTIFY_PORT_MAIL  = "6954"
)


var MailSrvNotify = NotifyObjectSt {
	addr:         NOTIFY_ADDR_SERV,
	port:         NOTIFY_PORT_AUTHSRV,
	serverMethod: "MailSrv.RecvCmd",
}

//END :-)

func (notify *NotifyObjectSt) SendMsg(cmd string, args interface{}, resp interface{}) bool {
	cli, err := rpc.DialHTTP("tcp", notify.addr+":"+notify.port)
	if err != nil {
		log.Println("dial notify failed: ", notify.serverMethod, err)
		return false
	}
	defer cli.Close()

	//package
	//var rep int
	var msg = NotifyMsgSt{Cmd: cmd, Args: args}
	err = cli.Call(notify.serverMethod, &msg, resp)
	if err != nil {
		log.Println("call failed: ", notify.serverMethod, err)
		return false
	}
	return true
}

func (notify *NotifyObjectSt) Services(recv interface{}) error {
	rpc.Register(recv)
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", notify.addr+":"+notify.port)
	if err != nil {
		log.Println("listen failed: ", notify.addr+":"+notify.port, err)
		return err
	}
	http.Serve(listen, nil)
	return nil
}
