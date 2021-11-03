package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// 产生log文件 分等级 debug
type Level int

const (
	DEBUG Level = iota
	INFO
	WARRING
	ERROR
	FATAL
)

var (
	F                  *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2 // 什么含义 没懂
	logger             *log.Logger
	logPrefix          = ""
	levelFlages        = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

func init() {
	// 文件夹的创建等
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func SetPrefix(log_level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlages[log_level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlages[log_level])
	}
	logger.SetPrefix(logPrefix)
}

func Debug(v ...interface{}) {
	SetPrefix(DEBUG)
	logger.Println(v)
}
func Info(v ...interface{}) {
	SetPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	SetPrefix(WARRING)
	logger.Println(v)
}
func Error(v ...interface{}) {
	SetPrefix(ERROR)
	logger.Println(v)
}
func Fatal(v ...interface{}) {
	SetPrefix(FATAL)
	logger.Println(v)
}
