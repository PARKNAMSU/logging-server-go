package util

import (
	"log"
	"time"

	"logFile.com/log-file-go/constant/localCache"
	"logFile.com/log-file-go/tool/common"
)

func RemoveTrashData() {
	loc, err := time.LoadLocation("Asia/Seoul")
	common.CheckErr(err)
	now := time.Now().In(loc).UnixMilli()

	if len(localCache.IdleProfFiles) == 0 {
		log.Println("profile data empty")
	}else {
		for key,data := range localCache.IdleProfFiles {
			if len(data) > 0 && now - data[0].CreatedTime > 6000 {
				delete(localCache.IdleProfFiles,key)
				log.Println("trash data removed")
			}
		}
	}
}
