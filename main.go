package main

import (
	"fmt"
	"github.com/greenpipig/get-GoCN-news/GoCN-news"
	"github.com/greenpipig/get-GoCN-news/getNews"
	"github.com/greenpipig/get-GoCN-news/log"
	"os/exec"
	"time"
)

func mainFunc() {
	url, title := getNews.FetchUrl("https://gocn.vip/topics/node18")
	log.Infof("newsPageUrl:%s,newsPageTitle:%s", url, title)
	if url == "" || title == "" {
		return
	}
	newsList, newsUrlList := getNews.FetchTotalNew(url)
	judgeFirst:=GoCN_news.WriteToMd(newsList, newsUrlList, title)
	if judgeFirst{
		command := `./update.sh`
		cmd := exec.Command("/bin/bash", "-c", command)
		_, err := cmd.Output()
		if err != nil {
			fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
			return
		}
	}
}

func main() {
	//go cron.ReloadJob("update mdFile", mainFunc, 1*time.Minute)
	for {
		mainFunc()
		time.Sleep(3 * time.Hour)
		//for test
		//time.Sleep(time.Minute)
	}
}
