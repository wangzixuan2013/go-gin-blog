package logging

import (
	"gin-blog/pkg/setting"
	"os"
	"time"
	"fmt"
	"log"
)

//var (
//	//LogSavePath = "runtime/logs/"
//	LogSavePath = setting.AppSetting.RuntimeRootPath + setting.AppSetting.LogSavePath
//	//LogSaveName = "log"
//	LogSaveName = setting.AppSetting.LogSaveName
//	//LogFileExt = "log"
//	LogFileExt = setting.AppSetting.LogFileExt
//	//TimeFormat = "20060102"
//	TimeFormat = setting.AppSetting.TimeFormat
//)

func getLogFilePath() string {
	dir, _ := os.Getwd()
	path := dir + "/" + setting.AppSetting.RuntimeRootPath + setting.AppSetting.LogSavePath
	return fmt.Sprintf("%s", path)
}

func getLogFileFullPath() string  {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt)
	return fmt.Sprintf("%s%s",prefixPath,suffixPath)
}

func getAccessLog() string{
	dir, _ := os.Getwd()
	prefixPath := dir + "/" + setting.AppSetting.RuntimeRootPath + "access_log/"
	suffixPath := fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt)
	return fmt.Sprintf("%s%s",prefixPath,suffixPath)
}

func openLogFile(filePath string) *os.File  {
	_, err := os.Stat(filePath)
	switch  {
		case os.IsNotExist(err):
			mkDir()
		case os.IsPermission(err):
			log.Fatalf("Permission :%v", err)
	}
	handle,err := os.OpenFile(filePath,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err !=nil{
		log.Println(filePath)
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	//log.Printf("openLogFile")
	return handle
}

func mkDir()  {
	dir,_ := os.Getwd()
	err := os.MkdirAll(dir + "/" + getLogFilePath(),os.ModePerm)
	if err != nil{
		panic(err)
	}
}










