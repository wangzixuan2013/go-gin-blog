package main

import (
	"gin-blog/models"
	"time"
	"log"
	"github.com/robfig/cron"
)

func main() {
	log.Println("Starting...")

	i := 0
	c := cron.New(cron.WithSeconds())
	spec := "*/1 * * * * *"
	c.AddFunc(spec, func() {
		models.CleanAllTag()
		i++
		log.Println("start", i)
	})
	c.Start()
	//select{} //阻塞主线程不退出
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
			print10()
		}
	}
}

func print5() {
	log.Println("Run 5s cron")
}

func print10() {
	log.Println("Run 10s cron")
}

func print15() {
	log.Println("Run 15s cron")
}
