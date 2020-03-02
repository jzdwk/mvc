/*
@Time : 20-3-2
@Author : jzd
@Project: mvc
*/
package models

import "time"

type Test struct {
	Id         int64      `orm:"column(id);pk;auto" json:"id,omitempty"`
	Name       string     `orm:"column(name);size(255)" json:"name"`
	CreateTime *time.Time `orm:"auto_now_add;type(datetime)" json:"createTime"`
	UpdateTime *time.Time `orm:"auto_now;type(datetime)" json:"updateTime"`
	Deleted    int8       `orm:"column(deleted);default(0)" json:"deleted"`
}

type testModel struct{}

func (t *Test) TableName() string {
	return "test"
}
