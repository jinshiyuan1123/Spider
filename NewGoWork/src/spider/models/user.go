package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Level     int       `json:"level"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UserToken string    `json:"user_token"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (self *User) TableName() string {
	return "user"
}

func UserAdd(user *User) (int64, error) {
	return orm.NewOrm().Insert(user)
}

func GetUserList(page, pageSize int, filters ...interface{}) ([]*User, int64) {
	offset := (page - 1) * pageSize
	list := make([]*User, 0)
	query := orm.NewOrm().QueryTable("user")
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("id").Limit(pageSize, offset).All(&list)
	return list, total
}

func (user *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(user, fields...); err != nil {
		return err
	}
	return nil
}

func GetUserById(id int) (*User, error) {
	user := new(User)
	err := orm.NewOrm().QueryTable(user.TableName()).Filter("id", id).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByName(name string) (*User, error) {
	user := new(User)
	err := orm.NewOrm().QueryTable(user.TableName()).Filter("username", name).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
