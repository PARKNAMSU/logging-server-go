package ui

import (
	commonConstant "logFile.com/log-file-go/constant/common"
	"logFile.com/log-file-go/controller/structure"
	"logFile.com/log-file-go/handler"
	"logFile.com/log-file-go/handler/profFileOpen"
	"logFile.com/log-file-go/tool/common"
)

func Controller(objectData structure.ClientData) {
	var handler handler.Handler
	switch objectData.Protocol {
	case commonConstant.PROTOCOL.ProfHttpOpen:
		var profOpenData profFileOpen.ProfFileOpenData
		common.SendDataJsonToStruct(objectData, &profOpenData)
		handler = &profOpenData
	}

	if handler != nil {
		handler.Execute()
	}
}
