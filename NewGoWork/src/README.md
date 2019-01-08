# 简单爬虫实现
使用golang实现简单的爬虫

使用cron定时爬取数据

## api接口

`/v1/book` : 获取小说列表

`/v1/book/getchapters?bookid=1&page=1` : 获取指定小说的章节，分页查询，每页10章

`/v1/book/chapter?id=950` : 获取指定章节详细内容
********************
## 上面都是原来作者的 想学习可以跟着看  

1.我改了新的爬虫方法因为原来的网页结构变了好像项目跑不起来  

2.改了一个小错误 原来的config文件有问题我把它改成了.ini文件  

3.我增加了部分注释给自己做笔记

## 这个项目想跑起来得需要以下部分
1.github.com/astaxie/beego 框架  
2. 其他的按照提示 go get就行了唯一一个个坑是有一个包要翻墙 不过https://www.golangtc.com/download/package 走这个网站下载就行了
