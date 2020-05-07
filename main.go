package main

import (
	"github.com/greenpipig/get-GoCN-news/GoCN-news"
	"github.com/greenpipig/get-GoCN-news/cron"
	"github.com/greenpipig/get-GoCN-news/getNews"
	"log"
	"os"
	"os/exec"
	"time"
)

var fileName = "GoCN-news.log"

func init(){
	f, err:= os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}
	log.SetPrefix("TRACE: ")
	log.SetOutput(f)
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
}

func mainFunc() {
	url, title := getNews.FetchUrl("https://gocn.vip/topics/node18")
	log.Printf("newsPageUrl:%s,newsPageTitle:%s", url, title)
	if url == "" || title == "" {
		return
	}
	newsList, newsUrlList := getNews.FetchTotalNew(url)
	judgeFirst := GoCN_news.WriteToMd(newsList, newsUrlList, title)
	if judgeFirst {
		command := `./update.sh`
		cmd := exec.Command("/bin/bash", "-c", command)
		_, err := cmd.Output()
		if err != nil {
			log.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
			return
		}
	}
}

func main() {
	cron.ReloadJob("update mdFile", mainFunc, 1*time.Hour)
}
