package getNews

import "testing"

func TestFetchTotalNew(t *testing.T) {
	FetchTotalNew("/topics/10337")
}

func TestFetchUrl(t *testing.T) {
	FetchUrl("https://gocn.vip/topics/node18")
}
