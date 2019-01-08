package models

import (
	"fmt"
	"time"
	"github.com/astaxie/beego/orm"
)

type Book struct{//不然返回的Key是大写的
	Id int `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	Image string `json:"image"`
	Status int `json:"status"`
	From string `json:"from"`
	Url string `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//`json:"updated_at"`
}

func BookAdd(book *Book)(int64, error){
	i, e := orm.NewOrm().Insert(book)
	return  i,e
}

func GetBookByName(name string)(*Book, error){
	book := new(Book)
	newOrm := orm.NewOrm()
	fmt.Println(newOrm)
	table := newOrm.QueryTable(&Book{})
	err := table.Filter("name", name).One(book)
	if err != nil || book.Id < 1{
		return nil, err
	}
	return book, nil
}

func GetBookList(filters ...interface{})([]*Book, int64){
	books := make([]*Book, 0)//这是个切片
	query := orm.NewOrm().QueryTable("book")
	if len(filters) > 0{
		l := len(filters)
        for i := 0; i < l; i += 2{
			query = query.Filter(filters[i].(string), filters[i + 1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("id").All(&books)
	return books, total
}