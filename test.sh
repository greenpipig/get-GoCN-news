time=$(date "+%Y-%m-%d")
title="-GOCN每日新闻.md"
git add .
git commit -m "update news"
git push origin master
# shellcheck disable=SC2164
cd GoCN-news
timeTitle=$time$title
cp $timeTitle ~/myblog/source/_posts/
# shellcheck disable=SC2164
cd ~/myblog #进入自己的blog地址
hexo clean
hexo d
