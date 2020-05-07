package getNews

import (
	"fmt"
	"github.com/greenpipig/soup"
	"log"
	//"github.com/anaskhan96/soup"
	"time"
)

func fetch(url string) soup.Root {
	fmt.Println("Fetch Url", url)
	soup.Headers = map[string]string{
		"User-Agent": "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
	}
	source, err := soup.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	doc := soup.HTMLParse(source)
	return doc
}

func FetchUrl(url string) (newsUrl string, title string) {
	doc := fetch(url)
	newsUrl = ""
	newsTitle := ""
	todayTimestamp := time.Now().Unix()                               //获得时间戳
	todayTimeStr := time.Unix(todayTimestamp, 0).Format("2006-01-02") //把时间戳转换成时间,并格式化为年月日
	wantNewTitleEn := "GOCN 每日新闻 (" + todayTimeStr + ")"
	wantNewTitleCN := "GOCN 每日新闻（" + todayTimeStr + "）"
	wantNewTitleEn1 := "GoCN 每日新闻 (" + todayTimeStr + ")"
	wantNewTitleCN1 := "GoCN 每日新闻（" + todayTimeStr + "）"
	for _, root := range doc.FindAllStrict("div", "class", "title media-heading") {
		newsUrl, _ = root.Find("a").Attrs()["href"]
		newsTitle, _ = root.Find("a").Attrs()["title"]
		message := root.Find("span", "class", "node").Text()
		fmt.Println(newsTitle)
		if message == "每日新闻" && (newsTitle == wantNewTitleEn || newsTitle == wantNewTitleCN || newsTitle == wantNewTitleEn1 || newsTitle == wantNewTitleCN1) {
			return newsUrl, newsTitle
		}
	}
	return "", ""
}

func FetchTotalNew(newsUrl string) (newsList []string, newsUrlList []string) {
	todayNewUrl := "https://gocn.vip" + newsUrl
	fmt.Println(todayNewUrl)
	doc := fetch(todayNewUrl)
	roots := doc.FindAllStrict("div", "class", "card-body markdown markdown-toc")
	root := roots[0].Find("ol")
	newsList = make([]string, 0, 10)
	newsUrlList = make([]string, 0, 10)
	for _, root1 := range root.FindAll("li") {
		err := ""
		func() {
			defer func() {
				if r := recover(); r != nil {
					err = "解析错误"
				}
			}()
			news := root1.Text()
			newsList = append(newsList, news)
			newsUrl := root1.Find("a", "rel", "nofollow").Text()
			newsUrlList = append(newsUrlList, newsUrl)
		}()
		if err != "" {
			if len(newsList)>len(newsUrlList){
				newsUrlList=append(newsUrlList,"")
			}
			continue
		}

	}
	return newsList, newsUrlList
}
