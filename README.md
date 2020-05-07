# get-GoCN-news
爬取获得gocn的新闻，并同步到博客，每日更新从我做起

具体部署方法：https://blog.csdn.net/weixin_44024220/article/details/105960728

### 参考库：

https://github.com/anaskhan96/soup

博客框架：hexo

### 特性

支持每天定时爬取

自动推送到hexo页面上

### 效果展示：

https://greenpipig.github.io/

### 使用方法：

首先修改update.sh中的路径文件，修改为自己的博客路径

每次检索时间为3小时一次

go build main.go

nohup ./main &

#### 踩坑：

对于html解析时该库无法解析空格，推荐使用doc.FindAllStrict此方法
