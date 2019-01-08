package spider

import (
	"errors"
	"fmt"
	"github.com/robfig/cron"
	"spider/config"
)

type SBook struct {
	Name     string
	Image    string
	Url      string
	Chapters []*SChapter
}

type SChapter struct {
	Title   string
	Url     string
	Order   int
	Pre     int
	Next    int
	Content string
}

type Spider interface {
	SpiderUrl(url string) error
}

func NewSpider(from string) (Spider, error) {
	switch from {
	case "booktxt":
		return new(BookTextSpider), nil
	default:
		return nil, errors.New("系统暂未处理该类型的配置文件")
	}
}

func Start() {

	c := cron.New()
	spec := config.AppConfig.GetString("task.spec")
	c.AddFunc(spec, getBook)
	c.Start()
}

func getBook() {
	fmt.Println("getBook()")
	s, err := NewSpider("booktxt")
	if err != nil {
		return
	}
	//http://www.booktxt.net/2_2219/
	go s.SpiderUrl("http://www.booktxt.net")
}
