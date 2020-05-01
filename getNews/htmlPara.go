package getNews

import (
	"fmt"
	"log"
	"src/github.com/anaskhan96/soup"
)

func Fetch(url string) soup.Root {
	fmt.Println("Fetch Url", url)
	soup.Headers = map[string]string{
		"User-Agent": "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
	}

	source, err := soup.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	doc := soup.HTMLParse(source)
	for _, root := range doc.FindAllStrict("div", "class", "title media-heading") {
		newsUrl, _ := root.Find("a").Attrs()["href"]
		movieUrl2, _ := root.Find("a").Attrs()["title"]
		title := root.Find("span", "class", "node").Text()
		fmt.Println(newsUrl, movieUrl2, title)
	}
	return doc
}



