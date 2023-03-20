package profFileOpen

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
	commonConstant "logFile.com/log-file-go/constant/common"
	"logFile.com/log-file-go/tool/awsModule"
	"logFile.com/log-file-go/tool/common"
	"logFile.com/log-file-go/tool/file"
)

var (
	_ = godotenv.Load()
)

func (profFileOpenData ProfFileOpenData) Execute() {
	output, _ := awsModule.GetObjectS3(
		profFileOpenData.Bucket,
		profFileOpenData.Location)
	bodyByte, err := ioutil.ReadAll(output.Body)
	if common.ErrorLogging(err) {
		return
	}

	log.Println("write")

	fileName := commonConstant.GetLogDirByEnv(os.Getenv("ENVIRONMENT")) + "/" +
		profFileOpenData.OpenType + "/" +
		"openHttp"

	file.WriteFile(fileName,
		bodyByte, 0777, commonConstant.FILE_EXTENSION.PROF)
	// shellModule.OpenProfHttp(profFileOpenData.OpenType, "6061")
}
