/*
@Time : 20-3-22
@Author : jzd
@Project: mvc
*/
package models

type Chunk struct {
	Id int64 `orm:"column(id);pk;auto" json:"id,omitempty"`
	//current num
	ChunkNumber int `orm:"column(chunk_number);size(255)" json:"chunkNumber"`
	//size
	ChunkSize        int64 `orm:"column(chunk_size)" json:"chunkSize"`
	CurrentChunkSize int64 `orm:"column(current_chunk_size)" json:"currentChunkSize"`
	TotalSize        int64 `orm:"column(total_size)" json:"totalSize"`
	//file uuid
	Identifier string `orm:"column(identifier)" json:"identifier"`
	FileName   string `orm:"column(file_name)" json:"fileName"`
	//save path
	RelativePath string `orm:"column(relative_path)" json:"relativePath"`
	TotalChunks  int    `orm:"column(total_chunks)" json:"totalChunks"`
	//file type
	Type string `orm:"column(type)" json:"type"`
}

type chunkModel struct{}

func (t *Chunk) TableName() string {
	return "chunk"
}

func (*chunkModel) GetChunksByIdentifier(identifier string) (rst []Chunk, err error) {
	qb := MysqlBuilder().Select(" * ").From(" chunk T0 ")
	qb.Where(" 1 = 1 ")
	var params []interface{}
	qb.And(" T0.Identifier = ? ")
	params = append(params, identifier)
	qb.OrderBy(" T0.ChunkNumber ASC ")
	_, err = Ormer().Raw(qb.String(), params).QueryRows(&rst)
	if err != nil {
		return nil, err
	}
	return
}
