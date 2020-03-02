package models

import (
	"github.com/astaxie/beego/orm"
	"sync"
)

var (
	globalOrm orm.Ormer
	once      sync.Once
	UserModel *userModel
)

func init() {
	//print sql
	orm.Debug = true
	// init orm tables
	orm.RegisterModel(new(User))
	// init models
	UserModel = &userModel{}
}

// singleton init ormer ,only use for normal db operation
// if you begin transaction，please use orm.NewOrm()
func Ormer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}
