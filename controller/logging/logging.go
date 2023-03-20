package logging

import (
	commonConstant "logFile.com/log-file-go/constant/common"
	"logFile.com/log-file-go/controller/structure"
	"logFile.com/log-file-go/handler"
	apilog "logFile.com/log-file-go/handler/apiLog"
	"logFile.com/log-file-go/handler/profFile"
	synclog "logFile.com/log-file-go/handler/syncLog"
	"logFile.com/log-file-go/handler/testLog"
	"logFile.com/log-file-go/tool/common"
)

func Controller(objectData structure.ClientData) {
	var handler handler.Handler
	switch objectData.Protocol {
	case commonConstant.PROTOCOL.SyncLog:
		var syncLogData synclog.SyncLogData
		common.SendDataJsonToStruct(objectData.Data, &syncLogData)
		handler = &syncLogData
	case commonConstant.PROTOCOL.TestLog:
		var testLogData testLog.TestLogData
		common.SendDataJsonToStruct(objectData.Data, &testLogData)
		handler = &testLogData
	case commonConstant.PROTOCOL.ApiLog:
		var apiLogData apilog.ApiLogData
		common.SendDataJsonToStruct(objectData.Data, &apiLogData)
		handler = &apiLogData
	case commonConstant.PROTOCOL.ProfFile:
		var profData profFile.ProfFileData
		common.SendDataJsonToStruct(objectData.Data, &profData)
		handler = &profData
	}

	if handler != nil {
		handler.Execute()
	}
}
