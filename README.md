# Local Logging Server

## Log 저장 위치

`$HOME/logs`

## GOLANG 세팅
```
1. go1.18 or go1.19 설치(홈페이지에서 패키지 설치 권장 - brew X) 설치페이지: https://golang.org/dl/
2. go env -w GO111MODULE=auto 설정
3. go mod download 의존성 패키지 설치
4. go get
```

## .env 세팅
```
TCP_PORT=":8000"
ENVIRONMENT="" // local or container
```

## Build 파일 생성
```shell
./script/go-build.sh
# 빌드파일 위치 ./go-build/log-file-go
```

## Docker Container 실행
```shell
# 먼저 .env -> ENVIRONMENT="container" 로 설정
./script/docker-run.sh
```

## 통신 프로토콜
```
TEST_LOG // 테스트용
SYNC_LOG // Batch SyncAll 성능 로그
API_LOG // API Request 로그
PROF_FILE // SyncAll profiling 전송

PROGRAM_EXIT // 프로그램 종료
```

## 서버 데이터 전달 포맷

```json
// sync log
{
    "protocol": "", // ... protocol,
    "data": {
        "alloc": 0, // 현재 힙 메모리 사용량
        "totalAlloc": 0, // 총 힙 메모리 사용량
        "sys": 0, // os 자원 사용량
        "numGC": 0, // GC 실행 횟수
        "runnigTime": 0, // 걸린 시간
    }
}

// sync profiling file send
{
    "protocol":"",
    "data": {
        "profType": "", // cpu(걸린시간) or mem(heap memory 사용량)
        "maxIdx": 0, // 파일 전송 최대 index (파일 크기가 크면 파일을 쪼개서 전송해야 함)
        "nowIdx": 0, // 파일 전송 현재 index
        "fileByte": [0,0], // 파일 바이트 배열
        "uuid": "" // 현재 전송중인 파일 파악 위한 uuid 키값
    }
}

// api log
{
    "protocol": "", // ... protocol,
    "data": {
        "body": "", // request body
        "header": "", // request header
        "url": "", // api url
        "ip": "", // client ip
        "serverType": "", // backend 타입 admin or user
        "error": "", // api 발생 에러
    }
}
```

