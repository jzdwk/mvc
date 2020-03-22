package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/mvc/models/upload"
	"sync"
)

var (
	globalOrm  orm.Ormer
	once       sync.Once
	UserModel  *userModel
	ChunkModel *upload.ChunkModel
	FileModel  *upload.FileInfoModel
)

func init() {
	//print sql
	orm.Debug = true
	// init orm tables
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Test))
	orm.RegisterModel(new(upload.Chunk))
	orm.RegisterModel(new(upload.FileInfo))
	// init models
	UserModel = &userModel{}
	ChunkModel = &upload.ChunkModel{}
	FileModel = &upload.FileInfoModel{}
}

// singleton init ormer ,only use for normal db operation
// if you begin transaction，please use orm.NewOrm()
func Ormer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}
