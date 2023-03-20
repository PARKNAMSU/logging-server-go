package system

import (
	"logFile.com/log-file-go/constant/common"
	"logFile.com/log-file-go/controller/structure"
)

func Controller(data structure.ClientData) bool {
	var endFlag bool = false

	switch data.Protocol {
	case common.PROTOCOL.ProgramExit:
		endFlag = true
	}
	return endFlag
}
