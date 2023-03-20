package common

import (
	"encoding/json"
	"log"
)

/*
	에러 체크 후 존재하면 시스템 종료
*/
func CheckErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

/*
	에러 체크 후 존재하면 로깅 후 true 반환
*/
func ErrorLogging(err error) bool {
	if err != nil {
		log.Printf("some error happend: %s", err.Error())
		return true
	}
	return false
}

/*
	전달받은 jsonData 를 struct 변수에 저장
*/
func SendDataJsonToStruct[V interface{}](jsonData interface{}, structData *V) {
	dataByte, err := json.Marshal(jsonData)
	CheckErr(err)
	err = json.Unmarshal(dataByte, structData)
	CheckErr(err)
}
