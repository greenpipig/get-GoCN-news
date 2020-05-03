package main

import (
	"github.com/greenpipig/get-GoCN-news/GoCN-news"
	"github.com/greenpipig/get-GoCN-news/getNews"
	"github.com/greenpipig/get-GoCN-news/log"
	"time"
)

func mainFunc() {
	url, title := getNews.FetchUrl("https://gocn.vip/topics/node18")
	log.Infof("newsPageUrl:%s,newsPageTitle:%s", url, title)
	if url == "" || title == "" {
		return
	}
	newsList, newsUrlList := getNews.FetchTotalNew(url)
	GoCN_news.WriteToMd(newsList, newsUrlList, title)
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
