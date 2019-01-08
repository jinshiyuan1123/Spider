package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"spider/common"
	"spider/models"
	"strconv"
	"strings"
	"sync"
	"time"
)

//div里边的id不能重复但是 class可以重复的 也就是这个div里边用过的class名字下一个div里边还能用
//这个地方是爬虫数据的地方 爬出来存在数据库的
type BookTextSpider struct {
}

// <ul> 这个标签是所有分类的书籍 只提供div级别的 #和div。方法其他都是直接 find（"标签名"）
//  class =r 最新入库小说
//<div class="l"> 最新更新小说
func (self *BookTextSpider) SpiderUrl(url string) error { //实现了 SpiderUrl 方法 爬取
	book := SBook{}
	book.Url = url
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}
	//这些注释是为了减少日志输出 这是基础代码 方笔以后自己看的
	doc.Find("#wrapper").Children().Each(func(i int, selection *goquery.Selection) {
		fmt.Println("wrapper doc.Find().Children().Each", i) //他有4个子节点

	})
	doc.Find("#wrapper").Each(func(i int, selection *goquery.Selection) {

		fmt.Println("wrapper doc.Find().Each", i) //他有1个子节点
	})
	doc.Find(".item").Each(func(i int, selection *goquery.Selection) { //这个是推荐的书
		author := selection.Find("span").Text() //因为当前就一个span 直接取
		href, _ := selection.Find("dt").Find("a").Attr("href")
		a := selection.Find("dt").Find("a").Text()
		fmt.Println(".item", i, "链接href=", common.GbkToUtf8(href), "书名=", common.GbkToUtf8(a), "作者=", common.GbkToUtf8(author)) //四个

	})

	//wrapper, _ := doc.Find("#wrapper").Html()//#是按div里边的id名字搜索节点
	//fmt.Println("wrapper=",wrapper)
	//footer, _ := doc.Find("div.footer").Html()//div.是按照div里边的class寻找
	//fmt.Println("footer=",footer)//
	//ywtop := doc.Find("div.ywtop").Text()
	//fmt.Println("text=",ywtop)

	main, _ := doc.Find("#main").Html()
	fmt.Println("main=", &main) //直接按照id查找就行了 这个查出来的就是整个的主要内容 其他都是顶部标题什么的|
	doc.Find("#main").Each(func(i int, selection *goquery.Selection) {
	})
	doc.Find("#main").Children().Each(func(i int, selection *goquery.Selection) { //find到的是自己 children才能遍历所有
	})
	chanfor_book := make(chan struct{})
	go func() {
		doc.Find(".novelslist").Each(func(i int, selection *goquery.Selection) { //这个地方插入书籍 详细内容得取到数据里里边的书籍的连接去取章节
			selection.Find("li").Each(func(i int, selection *goquery.Selection) {
				href, _ := selection.Find("a").Attr("href")
				bookname := selection.Find("a").Text()

				author := selection.Text()
				split := strings.Split(author, "/")                                                                                                                               //会多前面的内容 截取一下
				b := models.Book{Name: common.GbkToUtf8(bookname), CreatedAt: time.Now(), UpdatedAt: time.Now(), Author: common.GbkToUtf8(split[1]), Url: common.GbkToUtf8(href)} //先注释掉
				name, e := models.GetBookByName(b.Name)
				if e == nil || name == nil { //查询不到或者报错就插进去
					models.BookAdd(&b) //这个地方插入的时候 还是得把上面转为utf8 否则插不进去的
					//fmt.Println("models.BookAdd(&b)  err=", e1, "add=", add) //这个插入 因为name是 UNIQUE的 所以再次哈如也插入不进去的
				}
			})
		})
		close(chanfor_book)
	}()
	<-chanfor_book
	// at this point, all goroutines are ready to go - we just need to
	// tell them to start by closing the start channel slice 长度老是 0待解决
	var wg sync.WaitGroup
	all_book, _ := models.GetBookList()
	for _, book1 := range all_book {
		wg.Add(1)
		go func(singlebook *models.Book) {
			defer wg.Done()
			document, e := goquery.NewDocument(singlebook.Url)
			if e != nil {
				return
			}
			chapters := make([]models.Chapter, 0, 10)
			document.Find("dl").Children().Each(func(i int, selection *goquery.Selection) { //selection 这个selection里边有所有内容的

				content := selection.Find("a").Text()
				href, _ := selection.Find("a").Attr("href")
				chapter_title := common.GbkToUtf8(content)
				utf8href := common.GbkToUtf8(href)
				var checkcontent string

				var pre int
				if i > 1 {
					pre = i - 1
				} else {
					pre = 1
				}
				text := selection.Text()
				utf8 := common.GbkToUtf8(text)
				if len(chapter_title) > 0 {
					checkcontent = chapter_title
				} else {
					checkcontent = utf8
				}
				chapter := models.Chapter{BookId: singlebook.Id, Title: checkcontent, Content: singlebook.Url + utf8href, Sort: i, Pre: pre,
					Next: i + 1, UpdatedAt: time.Now(), CreatedAt: time.Now(), Bookidsort: strconv.Itoa(singlebook.Id) + "_" + singlebook.Url + utf8href}
				chapters = append(chapters, chapter)
			})
			for k, child := range chapters {
				if strings.Contains(child.Title, "正文") {
					chapter := chapters[k+1:]
					for _, chapter_s := range chapter {
						models.ChapterAdd(&chapter_s)
					}
					break
				}
			}
		}(book1)

	}
	wg.Wait()
	//这个章节爬出来是乱的
	// 开始爬章节了
	//	//分割线 下面是原来的代码
	//	bookname := common.GbkToUtf8(doc.Find("#info h1").Text())
	//
	//	b, err := models.GetBookByName(bookname)
	//	if err != nil || b == nil{
	//		b := models.Book{Name:bookname, CreatedAt:time.Now(), UpdatedAt:time.Now(),Url:"http://www.baidu.com"}
	//		models.BookAdd(&b)
	//	}书籍爬完了
	//	doc.Find("#list dd").Each(func (i int, contentSelection *goquery.Selection){
	//		if i < 9{
	//			return
	//		}
	//		pre := i - 9
	//		next := i -7
	//		title := common.GbkToUtf8(contentSelection.Find("a").Text())
	//		href, _ := contentSelection.Find("a").Attr("href")
	//		chapter := SChapter{Title:title,Url:"http://www.booktxt.net"+href, Order:i - 8, Pre:pre, Next:next}
	//		book.Chapters = append(book.Chapters, &chapter)
	//		u := models.Url{Url:chapter.Url}
	//		models.UrlAdd(&u)
	//	})
	////chapter表示 每本书籍的章节
	//	channel := make(chan struct{}, 100)//空内存的struct 不能呗输入
	//	for _, chapter := range book.Chapters{
	//		channel <- struct{}{} //这是一个匿名的struct struct{1}{2} 这块表示1表示 结构体的定义那个{} 2表示结构体初始化的内容发 不过是空的而已
	//		go SpiderChapter(b.Id, chapter, channel)
	//	}
	//
	//	for i := 0; i < 100; i++{//这个地方是等待任务结束的
	//		channel <- struct{}{}
	//	}
	//	close(channel)
	return nil
}

type ChanTag struct{}

func SpiderChapter(bookid int, chapter *SChapter, c chan struct{}) {
	defer func() { <-c }()
	if models.IsValidUrl(chapter.Url) {
		doc, err := goquery.NewDocument(chapter.Url)
		if err != nil {
			return
		}
		content := doc.Find("#content").Text()
		content = common.GbkToUtf8(content)
		content = strings.Replace(content, "聽", " ", -1)
		ch := models.Chapter{BookId: bookid, Title: chapter.Title, Content: content, Sort: chapter.Order, Pre: chapter.Pre, Next: chapter.Next, CreatedAt: time.Now(), UpdatedAt: time.Now()}
		models.ChapterAdd(&ch)
		models.SpideredUrl(chapter.Url)
	}
}
