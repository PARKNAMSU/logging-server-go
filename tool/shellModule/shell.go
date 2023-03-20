package shellModule

import (
	"log"
	"os/exec"

	"github.com/joho/godotenv"
	"logFile.com/log-file-go/tool/common"
)

var (
	_ = godotenv.Load()
)

func OpenCpuProfHttp() error {

	// cmdKilPort := exec.Command("fuser", "-k", "-n", "tcp", "6061")

	// cmdPprof := exec.Command("go", "tool", "pprof", "-http", ":6061", "/logging-batch-go/logs/cpu/openFile.prof")

	cmd := exec.Command("sh", "/logging-batch-go/script/open-cpu-prof.sh")

	runErr := cmd.Run()

	if common.ErrorLogging(runErr) {
		return runErr
	}

	// runErr = cmdPprof.Run()

	// if common.ErrorLogging(runErr) {
	// 	return runErr
	// }

	log.Println("exec success")

	return nil
}

func OpenMemProfHttp() error {

	cmd := exec.Command("sh", "/logging-batch-go/script/open-mem-prof.sh")

	runErr := cmd.Run()

	if common.ErrorLogging(runErr) {
		return runErr
	}
	return nil
}
