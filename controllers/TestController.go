/*
@Time : 20-3-2
@Author : jzd
@Project: mvc
*/
package controllers

import (
	"encoding/json"
	"github.com/mvc/controllers/base"
	"github.com/mvc/models"
	"github.com/mvc/token"
	"github.com/mvc/util/logs"
)

type TestController struct {
	base.ApiController
}

// Post ...
// @Title Post
// @Description create test
// @Param	body	body	models.Test	true	"body for test content"
// @Success 201 {string} success message
// @Failure 403 body is empty
// @router /  [post]
func (c *TestController) Post() {
	var v models.Test
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error("get body error. %v", err)
		c.AbortBadRequestFormat("test")
		return
	}
	if _, err := models.Add(new(models.Test), &v); err != nil {
		logs.Error("create test error. %v", err)
		c.HandleError(err)
		return
	}
	c.Success("create test success")
}

// GetOne Test
// @Title Get One
// @Description get method test
// @Param	id		path 	string	 true	"test id"
// @Success 200 {object} models.Test
// @Failure 403
// @router /:id  [get]
func (c *TestController) Get() {
	id := c.GetIDFromURL()
	test := models.Test{}
	err := models.GetOne(new(models.Test), &test, id)
	if err != nil {
		logs.Error("test error.", err)
		c.HandleError(err)
		return
	}
	c.Success(test)
}

// Put ...
// @Title Put
// @Description update the test
// @Param	id		path	string	true		"test id"
// @Param	body	body	models.Test	true	"body for test content"
// @Success 200 {string} success message
// @Failure 403 :id is not int
// @router /:id  [put]
func (c *TestController) Put() {
	id := c.GetIDFromURL()
	//get model by id
	v := models.Test{}
	err := models.GetOne(new(models.Test), &v, id)
	if err != nil {
		logs.Error("get test by id error.", err)
		c.HandleError(err)
		return
	}
	logs.Info("update test:", v.Name)
	//get model by body
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error("get request body error. %v", err)
		c.AbortBadRequestFormat("test")
		return
	}
	//reset fk
	if _, err := models.Update(new(models.Test), id, &v, []string{"Name", "UpdateTime"}); err != nil {
		logs.Error(" update test error. %v", err)
		c.HandleError(err)
	}
	c.Success("update test success ")
}

// Delete ...
// @Title Delete
// @Description delete the testt
// @Param	id		path	string	true	"test id"
// @Success 200 {string} delete success
// @Failure 403 id is empty
// @router /:id  [delete]
func (c *TestController) Delete() {
	id := c.GetIDFromURL()
	if _, err := models.SoftDelete(new(models.Test), id); err != nil {
		logs.Error(" delete test error. %v", err)
		c.HandleError(err)
	}
	c.Success("Delete test success")
}

// @router /upload [post]
func (c *TestController) Upload() {
	req := c.Ctx.Request
	p := c.ApiController.GetString("project")
	t := c.ApiController.GetString("tag")
	logs.Info(req.Header)
	logs.Info(req.Body)
	logs.Info(req.URL)
	c.Success("test" + p + t)
}

// @Title Auth
// @Description docker v2 registry 3rd-part authNZ service
// @Failure 400
// @router /auth  [get]
func (c *TestController) Auth() {

	request := c.Ctx.Request
	logs.Info("URL for token request: %s", request.URL.String())
	//get harbor-registry
	service := c.GetString("service")
	//user
	usr := c.GetString("account")
	tokenCreator := token.GeneralCreator{Service: service, Usr: usr}
	token2, err := tokenCreator.Create(request)
	if err != nil {
		logs.Error(err.Error())
	}
	c.Data["json"] = token2
	c.ServeJSON()
}
