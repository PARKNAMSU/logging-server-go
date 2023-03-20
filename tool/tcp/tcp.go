package tcp

import (
	"net"

	"logFile.com/log-file-go/tool/common"
)

/*
	설정에 맞는 network 연결하여 연결된 리스너 리턴
*/
func GetListener(method string, port string) net.Listener {
	l, err := net.Listen(method, port)
	common.CheckErr(err)
	return l
}

/*
	리스너 통해 연결된 Connect 리턴
*/
func AcceptConnection(l net.Listener) net.Conn {
	conn, err := l.Accept()
	common.CheckErr(err)
	return conn
}