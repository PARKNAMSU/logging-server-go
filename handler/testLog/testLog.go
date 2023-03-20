package testLog

import (
	"strconv"
	"time"

	commonConstant "logFile.com/log-file-go/constant/common"
	"logFile.com/log-file-go/tool/common"
	"logFile.com/log-file-go/tool/file"
)

func (data *TestLogData) Execute() {
	loc, err := time.LoadLocation("Asia/Seoul")
	common.CheckErr(err)

	now := time.Now().In(loc)

	fileName := commonConstant.
		GetLogDirByEnv(commonConstant.GetEnvironment()) +
		"/test/" +
		now.Format("2006-01-02")

	if !file.FileExistCheck(fileName, commonConstant.FILE_EXTENSION.CSV) {
		file.WriteCSVFile(fileName,
			[]string{"ID", "CODE"}, 0777)
	}
	file.WriteCSVFile(fileName,
		[]string{strconv.Itoa(data.Id), string(data.Code)}, 0777)
}
