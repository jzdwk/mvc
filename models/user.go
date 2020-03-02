package models

import (
	"mvc/common"
	"mvc/util/encode"
	"time"
)

type UserType int

const PWD_SALT = "cmCc#1.sgm&168"

type userModel struct{}

type User struct {
	Id   int64  `orm:"pk;auto" json:"id"`
	Name string `orm:"index;size(200)" json:"name"`

	Password string `orm:"size(255)" json:"password"`
	Email    string `orm:"size(200)" json:"email"`
	Phone    string `orm:"size(200)" json:"phone"`
	Admin    bool   `orm:"default(False)" json:"admin"`

	LastLogin  string     `orm:"size(200)" json:"lastLogin"`
	LastIp     string     `orm:"size(200)" json:"lastIp"`
	Deleted    bool       `orm:"default(false)" json:"deleted"`
	CreateTime *time.Time `orm:"auto_now_add;type(datetime)" json:"createTime"`
	UpdateTime *time.Time `orm:"auto_now;type(datetime)" json:"updateTime"`
}

type UserView struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`

	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Admin      bool   `json:"admin"`
	LastLogin  string `json:"lastLogin"`
	LastIp     string `json:"lastIp"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type UserStatistics struct {
	Total int64 `json:"total,omitempty"`
}

func (*User) TableName() string {
	return "user"
}

func (*userModel) GetUserDetail(name string) (user *User, err error) {
	user1 := &User{}
	qs := Ormer().QueryTable("user").Filter("Deleted", 0).Filter("Name", name)
	if err := qs.One(user1); err != nil {
		return nil, err
	}
	return user1, nil
}

func (*userModel) UpdateUserAdmin(m *User) (err error) {
	v := &User{Id: m.Id}
	if err = Ormer().Read(v); err != nil {
		return
	}
	v.Admin = m.Admin
	_, err = Ormer().Update(v)
	return
}

func (*userModel) ResetUserPassword(id int64, password string) (err error) {
	v := &User{Id: id}
	if err = Ormer().Read(v); err != nil {
		return
	}
	salt := encode.GetRandomString(10)
	passwordHashed := encode.EncodePassword(password, salt)

	v.Password = passwordHashed
	_, err = Ormer().Update(v)
	return
}

func (*userModel) UpdateUserById(m *User) (err error) {
	v := &User{Id: m.Id}
	if err = Ormer().Read(v); err != nil {
		return
	}
	if m.Password != "" {
		v.Password = encode.EncodePassword(m.Password, PWD_SALT)
	}

	if m.Email != "" {
		v.Email = m.Email
	}
	if m.Phone != "" {
		v.Phone = m.Phone
	}

	v.Admin = m.Admin
	_, err = Ormer().Update(v)
	return

}

func (*userModel) GetUserCount4View(q *common.QueryParam) (total int64, err error) {
	qb := mysqlBuilder().Select("count(T0.id)").From("user T0")
	qb.Where(" T0.deleted = 0")
	var params []interface{}
	if name := q.Query["name"]; name != nil && name != "" {
		qb.And(" T0.name = ? ")
		params = append(params, q.Query["name"])
	}
	err = Ormer().Raw(qb.String(), params).QueryRow(&total)
	return
}

func (*userModel) GetUserPaged4View(q *common.QueryParam) (rst []UserView, err error) {
	qb := mysqlBuilder().Select("T0.id,T0.name,T0.email,T0.phone,T0.admin,T0.last_login,T0.last_ip,T0.create_time,T0.update_time").From("user T0")
	qb.Where(" T0.deleted = 0 ")
	var params []interface{}
	if name := q.Query["name"]; name != nil && name != "" {
		qb.And(" T0.name = ? ")
		params = append(params, q.Query["name"])
	}
	qb = BuildGroupBy(qb, q.GroupBy)
	qb = BuildOrder(qb, q.Order)
	qb = qb.Limit(int(q.Limit())).Offset(int(q.Offset()))
	_, err = Ormer().Raw(qb.String(), params).QueryRows(&rst)
	if err != nil {
		return nil, err
	}
	return
}

func (*userModel) GetAll(users *[]*User) (num int64, err error) {

	if num, err := Ormer().QueryTable(&User{}).Filter("Deleted", 0).All(users, "Id", "Name"); err == nil {
		return num, nil
	}
	return -1, err
}
