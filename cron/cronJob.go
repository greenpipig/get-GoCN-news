package cron

import (
	"log"
	"math/rand"
	"time"
)

func ReloadJob(jobName string, reloadFunc func(), intervalTime time.Duration) {
	defer func() {
		log.Printf("[cron_job]%s,fail", jobName)
		recover()
	}()
	// 先随机sleep一段时间, [0, 1000)毫秒
	randMillisecond := rand.Int63n(1000)
	time.Sleep(time.Duration(randMillisecond) * time.Millisecond)
	for {
		reloadFunc()
		time.Sleep(intervalTime)
	}
}
