# [Terms of Service Knight](http://alpha.tosknight.org/)

[![Go Report Card](https://goreportcard.com/badge/github.com/siglt/tosknight)](https://goreportcard.com/report/github.com/siglt/tosknight)

## 计划内功能

* 定期爬取腾讯、网易、知乎、百度等网站的用户协议（ http://www.qq.com/contract.shtml ）、隐私条款；
* 如果较之先前的版本有变化，保存新版的协议内容，并对变化内容进行比较；
* 可以添加需要跟踪的网站页面；
* 可能会考虑链接可信时间戳的API。

## 实现思路

使用 [html2text](https://github.com/Alir3z4/html2text/) 输出内容，使用 Unix 内置 diff 判断文件修改，使用 [diff2html](https://github.com/rtfpessoa/diff2html) 渲染输出。

## 相关项目

* https://tosdr.org/index.html https://github.com/tosdr/tosdr.org
* https://github.com/ecprice/newsdiffs

## LICENSE

GPLv3
