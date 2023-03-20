package profFile

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/novalagung/gubrak/v2"
	"logFile.com/log-file-go/constant/localCache"
	"logFile.com/log-file-go/structures"
	"logFile.com/log-file-go/tool/awsModule"
	"logFile.com/log-file-go/tool/common"
)

// var (
// 	IdleProfFiles = map[string][]OrderFile{}
// )
var (
	_ = godotenv.Load()
)

func (profData ProfFileData) Execute() {
	loc, err := time.LoadLocation("Asia/Seoul")
	common.CheckErr(err)

	now := time.Now().In(loc)
	// fileName := commonConstant.GetLogDirByEnv(commonConstant.GetEnvironment()) + "/prof/" + now.Format("2006-01-02")

	if localCache.IdleProfFiles[profData.Uuid] == nil {
		localCache.IdleProfFiles[profData.Uuid] = []structures.OrderFile{
			{
				File:        profData.FileByte,
				Order:       profData.NowIdx,
				CreatedTime: now.UnixMilli(),
			},
		}
	} else {
		localCache.IdleProfFiles[profData.Uuid] = append(localCache.IdleProfFiles[profData.Uuid], structures.OrderFile{
			File:  profData.FileByte,
			Order: profData.NowIdx,
		})
	}

	if profData.NowIdx >= profData.MaxIdx {

		makeFile := []byte{}

		orderData := gubrak.From(localCache.IdleProfFiles[profData.Uuid]).OrderBy(func(orderb structures.OrderFile) int {
			return orderb.Order
		}).Result().([]structures.OrderFile)

		gubrak.From(orderData).Each(func(data structures.OrderFile) {
			makeFile = append(makeFile, data.File...)
		})

		log.Println("write prof")

		result, err := awsModule.UploadS3(
			os.Getenv("LOGGING_BUCKET"),
			"profiling/"+
				profData.ProfType+"/"+
				now.Format("2006-01-02")+"/"+
				profData.Uuid+
				"-"+profData.ProfType+
				".prof", string(makeFile), "multipart/formed-data")
		log.Printf("%v , %v\n", result, err)
		delete(localCache.IdleProfFiles, profData.Uuid)

		// file.WriteFile(
		// 	commonConstant.GetLogDirByEnv(os.Getenv("ENVIRONMENT"))+"/"+
		// 		profData.ProfType+"/"+
		// 		"openHttp",
		// 	makeFile,
		// 	755,
		// 	commonConstant.FILE_EXTENSION.PROF)

		// var shellError error
		// if profData.ProfType == "cpu" {
		// 	shellError = shellModule.OpenCpuProfHttp()
		// } else {
		// 	shellError = shellModule.OpenMemProfHttp()
		// }
		// common.ErrorLogging(shellError)
	}
}
