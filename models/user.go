package models

import (
	"mvc/util/encode"
	"time"
)

type UserType int

const PWD_SALT = "mvc#1.&168"

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

func (*userModel) GetAll(users *[]*User) (num int64, err error) {

	if num, err := Ormer().QueryTable(&User{}).Filter("Deleted", 0).All(users, "Id", "Name"); err == nil {
		return num, nil
	}
	return -1, err
}
