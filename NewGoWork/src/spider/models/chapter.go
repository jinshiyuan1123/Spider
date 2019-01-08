package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Chapter struct{
	Id int `json:"id"`
	BookId int `json:"book_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Sort int `json:"sort"`//这个字段是用来排序的
	Pre int `json:"pre"`
	Next int `json:"next"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Bookidsort string `json:"bookidsort"`
} 

func ChapterAdd(chapter *Chapter)(int64, error){
	i, e := orm.NewOrm().Insert(chapter)
	//fmt.Println("ChapterAdd i=",i,"e=",e)
	return i,e
}

func GetChapterPage(page, pageSize int, filters ...interface{})([]*Chapter, int64){
	offset := (page - 1) * pageSize
	list := make([]*Chapter, 0)
	query := orm.NewOrm().QueryTable("chapter")
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("sort").Limit(pageSize, offset).All(&list)
	return list, total
}

	func GetChapterById(id int)(*Chapter, error){
	chapter := new(Chapter)
	err := orm.NewOrm().QueryTable("chapter").Filter("id", id).One(chapter)
	if err != nil{
		return nil, err
	}
	return chapter, nil
}