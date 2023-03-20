package synclog

import (
	"strconv"
	"time"

	commonConstant "logFile.com/log-file-go/constant/common"
	"logFile.com/log-file-go/tool/common"
	"logFile.com/log-file-go/tool/file"
)

func (data *SyncLogData) Execute() {
	loc, err := time.LoadLocation("Asia/Seoul")
	common.CheckErr(err)

	now := time.Now().In(loc)

	var branch string = data.Branch

	if data.Branch == "" {
		branch = "master"
	}

	fileName := commonConstant.
		GetLogDirByEnv(commonConstant.GetEnvironment()) +
		"/sync/" +
		branch + "_" + now.Format("2006-01-02")

	if !file.FileExistCheck(fileName, commonConstant.FILE_EXTENSION.CSV) {
		file.WriteCSVFile(fileName,
			[]string{"Date Time", "Alloc(MB)", "Total Alloc(MB)", "Sys(MB)", "NumGC(count)", "Runnig Time(s)"}, 0777)
	}
	file.WriteCSVFile(fileName,
		[]string{
			now.Format("2006-01-02T15:04:05"),
			strconv.Itoa(data.Alloc),
			strconv.Itoa(data.TotalAlloc),
			strconv.Itoa(data.Sys),
			strconv.Itoa(data.NumGC),
			strconv.Itoa(data.RunningTime)}, 0777)
}
