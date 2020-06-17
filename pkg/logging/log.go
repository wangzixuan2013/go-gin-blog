package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix = ""
	DefaultCallerDepth = 2

	logger *log.Logger
	logPrefix = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

	//fileName = ""
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)


func Setup()  {
	var err error
	//filePath := getLogFileFullPath()
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = openLogFile(fileName,filePath)
	if err != nil {
		log.Fatalln(err)
	}
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func AccessLog() string  {

	return getAccessLog()
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {

	setPrefix(INFO)

	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	fileFullPath := filePath + fileName
	//每天一个log
	_, err = os.Stat(fileFullPath)
	if err != nil{
		filePath := getLogFilePath()
		fileName := getLogFileName()
		F, err = openLogFile(fileName,filePath)
		if err != nil {
			log.Fatalln(err)
		}
		logger = log.New(F, DefaultPrefix, log.LstdFlags)
	}
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}



