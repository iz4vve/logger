package logger

import (
	"fmt"
	"path"
	"runtime"

	"github.com/hhkbp2/go-logging"
)

func GetCustomLogger() logging.Logger {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Failed to load config.yaml")
	}
	config := path.Join(path.Dir(filename), "../logger/config.yaml")
	if err := logging.ApplyConfigFile(config); err != nil {
		panic(err.Error())
	}
	logger := logging.GetLogger("main")
	return logger
}
