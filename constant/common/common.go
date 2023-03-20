package common

import (
	"os"

	"github.com/joho/godotenv"
)

type Protocol struct {
	// logging
	TestLog  string
	SyncLog  string
	ApiLog   string
	ProfFile string

	//ui
	ProfHttpOpen string

	ProgramExit string
}

type FileExtension struct {
	TEXT string
	CSV  string
	PROF string
}

type FileFlags struct {
	CRW_APPEND int
	CRW_TRUNC  int
	C_SYNC     int
}


const (
	// 서버 연결 방식
	METHOD = "tcp"
)

var (
	_ = godotenv.Load()
)


var (
	SAVE_MODE_UNITS = map[string]string{
		"AWS": "aws",
		"LOCAL": "local",
	}
	// 로그 저장 방식 
	SAVE_MODE = SAVE_MODE_UNITS["AWS"]

	// 프로그램 프로토콜
	PROTOCOL = Protocol{
		// logging
		TestLog:  "TEST_LOG",
		SyncLog:  "SYNC_LOG",
		ApiLog:   "API_LOG",
		ProfFile: "PROF_FILE",

		//ui
		ProfHttpOpen: "PROF_HTTP_OPEN",

		// system
		ProgramExit: "PROGRAM_EXIT",
	}

	// 로깅 관련 프로토콜
	LOGGING_PROTOCOLS = [3]string{
		PROTOCOL.SyncLog,
		PROTOCOL.ApiLog,
		PROTOCOL.ProfFile,
	}

	// UI 오픈 프로토콜
	UI_PROTOCOLS = [1]string{
		PROTOCOL.ProfHttpOpen,
	}

	// 시스템 설정 관련 프로토콜
	SYSTEM_PROTOCOLS = [1]string{
		PROTOCOL.ProgramExit,
	}

	// 파일 확장자
	FILE_EXTENSION = FileExtension{
		CSV:  ".csv",
		TEXT: ".txt",
		PROF: ".prof",
	}

	// 파일 플래그
	FILE_FLAGS = FileFlags{
		CRW_APPEND: os.O_CREATE | os.O_RDWR | os.O_APPEND,
		CRW_TRUNC:  os.O_CREATE | os.O_RDWR | os.O_TRUNC,
		C_SYNC:     os.O_CREATE | os.O_SYNC,
	}
)

/*
	서버 실행중인 유저 홈 디렉토리 경로
*/
func GetUserHomePath() string {
	home, _ := os.UserHomeDir()
	return home
}

/*
	실행 환경 (local, container)
*/
func GetEnvironment() string {
	var env string = os.Getenv("ENVIRONMENT")
	if env == "" {
		return "local"
	}
	return env
}

/*
	환경변수에 따른 로그 저장 경로
*/
func GetLogDirByEnv(env string) string {
	switch env {
	case "container":
		return "/logging-batch-go/logs"
	case "local":
		return GetUserHomePath() + "/logs"
	default:
		return GetUserHomePath() + "/logs"
	}
}

/*
	tcp 연결 포트
*/
func GetPort() string {
	var port = os.Getenv("TCP_PORT")
	if port == "" {
		return ":8000"
	}
	return port
}
