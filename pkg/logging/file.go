package logging

import (
	"gin-blog/pkg/file"
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

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
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

func openLogFile(fileName,filePath string) (*os.File,error)  {

	src := filePath
	perm := file.CheckPermission(src)

	//fmt.Println(src)
	//fmt.Println(fileName)

	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err := file.IsNotExistMkDir(src)

	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	handle,err := file.Open(src + fileName,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err !=nil{
		log.Println(src)
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}
	//log.Printf("openLogFile")
	return handle,nil
}











