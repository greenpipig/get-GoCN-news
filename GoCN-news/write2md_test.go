package GoCN_news

import (
	"src/github.com/stretchr/testify/assert"
	"testing"
)

func TestPathExists(t *testing.T) {
	exist, _ := PathExists("2020-05-02-GOCN每日新闻.md")
	assert.Equal(t, true, exist)
}
