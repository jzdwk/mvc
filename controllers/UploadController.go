/*
@Time : 20-3-22
@Author : jzd
@Project: mvc
*/
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/mvc/controllers/base"
	cupload "github.com/mvc/controllers/upload"
	"github.com/mvc/models"
	"github.com/mvc/models/upload"
	"github.com/mvc/util/logs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type UploadController struct {
	base.ApiController
}

const FILE_PATH = "/home/jzd/tmp"

// Post ...
// @Title Post
// @Description create test
// @Param	body	body	upload.Chunk	true	"body for test content"
// @Success 201 {string} success message
// @Failure 403 body is empty
// @router /  [post]
func (c *UploadController) Post() {
	//get chunk and info
	chunk, _, err := c.GetFile("upload")
	if err != nil {
		logs.Error("get file error. %v", err)
		c.AbortBadRequestFormat("get upload file err.")
		return
	}

	var v upload.Chunk
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error("get body error. %v", err)
		c.AbortBadRequestFormat("test")
		return
	}

	//file operate   "tmp/345345/1"
	path := FILE_PATH + "/" + v.Identifier + "/"
	fileName := strconv.Itoa(v.ChunkNumber)
	content, err := ioutil.ReadAll(chunk)
	if err := SaveFile(content, path+fileName); err != nil {
		logs.Error("chunk save error. %v", err)
		c.AbortBadRequestFormat("chunk save error. ")
		return
	}

	if _, err := models.Add(new(upload.Chunk), &v); err != nil {
		logs.Error("create test error. %v", err)
		c.HandleError(err)
		return
	}
	c.Success("chunk post success")
}

// @Title CheckChunkInfo
// @Description  test
// @Param	chunknum	param	int	true	"chunk num"
// @Param	identifier	param	string	true	"file id"
// @Success 201 {string} success message
// @Failure 403 body is empty
// @router /  [get]
func (c *UploadController) CheckChunk() {
	chunkNum, _ := c.GetInt("chunknum")
	identifier := c.GetString("identifier")
	param := make(map[string]interface{})
	param["Identifier"] = identifier
	param["ChunkNumber"] = chunkNum
	if models.IsExist(upload.Chunk{}, param) {
		c.AbortBadRequest("chunk already exist")
	}
	c.Success("chunk ok")
}

// Post ...
// @Title Post
// @Description create test
// @Param	body	body	upload.FileInfo	true	"body for test content"
// @Success 201 {string} success message
// @Failure 403 body is empty
// @router /  [post]
func (c *UploadController) MergeFile() {
	//get file info
	var v upload.FileInfo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error("get body error. %v", err)
		c.AbortBadRequestFormat("test")
		return
	}
	filename := v.FileName
	file := FILE_PATH + "/" + v.Identifier + "/" + filename
	if err := cupload.Merge(file, v.Identifier); err != nil {
		logs.Error("merge file error. %v", err)
		c.AbortBadRequest("merge file error.")
	}
	c.Success("merge success")
}

func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func MakeDir(dir string) error {
	if !FileIsExisted(dir) {
		if err := os.MkdirAll(dir, 0777); err != nil { //os.ModePerm
			fmt.Println("MakeDir failed:", err)
			return err
		}
	}
	return nil
}

func SaveFile(content []byte, path string) error {
	dir := filepath.Dir(path)
	if !FileIsExisted(dir) {
		MakeDir(dir)
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(content)
	if err != nil {
		return err
	}
	return nil
}
