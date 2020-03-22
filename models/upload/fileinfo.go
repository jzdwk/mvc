/*
@Time : 20-3-22
@Author : jzd
@Project: mvc
*/
package upload

type FileInfo struct {
	Id         int64  `orm:"column(id);pk;auto" json:"id,omitempty"`
	FileName   string `orm:"column(filename);size(255)" json:"fileName"`
	Identifier string `orm:"column(identifier)" json:"Identifier"`
	TotalSize  int64  `orm:"column(total_size)" json:"totalSize"`
	Type       string `orm:"column(type)" json:"type"`
	Location   string `orm:"column(location)" json:"location"`
}

type FileInfoModel struct{}

func (t *FileInfo) TableName() string {
	return "fileinfo"
}
