package main

import (
	"fmt"
	"gin-blog/models"
	"gin-blog/pkg/logging"
	"github.com/fvbock/endless"
	"log"
	//"net/http"
	"syscall"

	"gin-blog/routers"

	"gin-blog/pkg/setting"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	log.Println(setting.ServerSetting.ReadTimeout)
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		logging.Info("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}

