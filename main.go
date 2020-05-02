package main

import (
	"get-GoCN-news/GoCN-news"
	"get-GoCN-news/getNews"
	"get-GoCN-news/log"
)

func main() {
	url, title := getNews.FetchUrl("https://gocn.vip/topics/node18")
	log.Infof("newsPageUrl:%s,newsPageTitle:%s", url, title)
	if url == "" || title == "" {
		return
	}
	newsList, newsUrlList := getNews.FetchTotalNew(url)
	GoCN_news.WriteToMd(newsList, newsUrlList,title)
}
