package main

import (
	"log"
	"runtime"
	"time"

	"logFile.com/log-file-go/constant/common"
	"logFile.com/log-file-go/controller"
	"logFile.com/log-file-go/tool/tcp"
	"logFile.com/log-file-go/tool/util"
)

var (
	END_FLAG = false
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	l := tcp.GetListener(common.METHOD, common.GetPort())

	defer l.Close()
	defer log.Println("process exit!!!!")

	go func() {
		for !END_FLAG {
			util.RemoveTrashData()
			time.Sleep(time.Second * 60)
		}
		log.Println("remove trash data end")
	}()

	for !END_FLAG {
		if controller.Controller(l) {
			END_FLAG = true
		}
		continue
	}

}
