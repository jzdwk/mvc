package models

import (
	"github.com/astaxie/beego/orm"
	"sync"
)

var (
	globalOrm  orm.Ormer
	once       sync.Once
	UserModel  *userModel
	ChunkModel *chunkModel
	FileModel  *fileInfoModel
)

func init() {
	//print sql
	orm.Debug = true
	// init orm tables
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Test))
	orm.RegisterModel(new(Chunk))
	orm.RegisterModel(new(FileInfo))
	// init models
	UserModel = &userModel{}
	ChunkModel = &chunkModel{}
	FileModel = &fileInfoModel{}
}

// singleton init ormer ,only use for normal db operation
// if you begin transaction，please use orm.NewOrm()
func Ormer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}
