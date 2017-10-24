# Terms of Service Knight

## 实现思路

利用 Git 来做内容的 diff，如果有修改就 dump 成文件提交到 storage repo 中，原文件和 diff 输出分别存储，原文件用来做 diff 和备份，diff 用来前端展示。通知可以考虑 web hook，但是这样需要一个服务器，有支出。最好是发一个 PR，@ 感兴趣的所有人，让 GitHub 帮忙发邮件。

## 进度

### TODO

* Config 文件读取有问题，考虑是不是直接用 cobra 集成的
* 内容 diff 和 modified 的实现与测试

## 计划内功能

* 定期爬取腾讯、网易、知乎、百度等网站的用户协议（ http://www.qq.com/contract.shtml ）、隐私条款；
* 如果较之先前的版本有变化，保存新版的协议内容，并对变化内容进行比较；
* 可以添加需要跟踪的网站页面；
* 可能会考虑链接可信时间戳的API。

## 相关项目

* https://tosdr.org/index.html https://github.com/tosdr/tosdr.org
* https://github.com/ecprice/newsdiffs

## LICENSE

GPLv3
