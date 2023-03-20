package controller

import (
	"encoding/json"
	"log"
	"net"

	"github.com/novalagung/gubrak/v2"
	commonConstant "logFile.com/log-file-go/constant/common"
	"logFile.com/log-file-go/controller/logging"
	"logFile.com/log-file-go/controller/structure"
	"logFile.com/log-file-go/controller/system"
	"logFile.com/log-file-go/tool/common"
	"logFile.com/log-file-go/tool/tcp"
)

func Controller(l net.Listener) bool {
	var conn net.Conn = tcp.AcceptConnection(l)
	var endFlag bool = false
	var recvBuf []byte = make([]byte, 4096)
	n, err := conn.Read(recvBuf)
	if common.ErrorLogging(err) && n == 0 {
		if err != nil {
			log.Printf("not accepted tcp connect [recived buffer]:[%d] , [error]:[%s]\n", n, err.Error())
		} else {
			log.Printf("not accepted tcp connect [recived buffer]:[%d]\n", n)
		}
		return endFlag
	}

	var objectData structure.ClientData
	data := recvBuf[:n]
	err = json.Unmarshal(data, &objectData)

	if err != nil {
		log.Println(err)
		conn.Close()
		return endFlag
	}

	switch {
	case gubrak.
		From(commonConstant.LOGGING_PROTOCOLS).
		IndexOf(objectData.Protocol).Result() >= 0:
		logging.Controller(objectData)
	case gubrak.
		From(commonConstant.UI_PROTOCOLS).
		IndexOf(objectData.Protocol).Result() >= 0:
	case gubrak.
		From(commonConstant.SYSTEM_PROTOCOLS).
		IndexOf(objectData.Protocol).
		Result() >= 0:
		endFlag = system.Controller(objectData)
	default:
		log.Println("unusable protocol")
	}
	conn.Close()
	return endFlag
}
