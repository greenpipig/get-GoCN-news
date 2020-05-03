time=$(date "+%Y-%m-%d")
title="-GOCN每日新闻.md"
git add GoCN-news
git commit -m "update news"
git push origin master
# shellcheck disable=SC2164
cd GoCN-news
timeTitle=$time$title
cp $timeTitle ~/myblog/source/_posts/#将生成的md文件放置到hexo目录下，此处要修改路径名
# shellcheck disable=SC2164
cd ~/myblog #进入自己的blog地址，此处要修改路径名
hexo clean
hexo d
