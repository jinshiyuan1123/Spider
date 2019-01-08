package models

import(
	"github.com/astaxie/beego/orm"
)

type Url struct{
	Id int `json:"id"`
	Url string `json:"url"`
	Status int `json:"status"`
}

func UrlAdd(url *Url)(int64, error){
	return orm.NewOrm().Insert(url)
}

func GetUrlByUrl(url string)(*Url, error){
	u := new(Url)
	err := orm.NewOrm().QueryTable("url").Filter("url", url).One(u)
	if err != nil{
		return nil, err
	}
	return u, nil
}

func IsValidUrl(url string)bool{
	u, err := GetUrlByUrl(url)
	if err != nil{
		return false
	}
	if u.Status == 0{
		return true
	}
	return false
}

func (self *Url)Unavailable()bool{
	url, err := GetUrlByUrl(self.Url)
	if err != nil{
		return false
	}
	url.Status = -1
	_, err = orm.NewOrm().Update(url, "status")
	if err != nil{
		return false
	}
	return true
}

func SpideredUrl(url string)bool{
	u, err := GetUrlByUrl(url)
	if err != nil{
		return false
	}
	u.Status = 1
	_, err = orm.NewOrm().Update(u, "status")
	if err != nil{
		return false
	}
	return true
}