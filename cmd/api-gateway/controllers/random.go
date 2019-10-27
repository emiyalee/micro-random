package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/emiyalee/micro-random/cmd/api-gateway/models"
)

// RandomController operations for Random
type RandomController struct {
	beego.Controller
	Client    *models.RandomClient
	RandomNum int32
}

// URLMapping ...
func (c *RandomController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

type randomResponse struct {
	Error     string `json:"error_code"`
	RandomNum int32  `json:"random_num"`
}

// Post ...
// @Title Create
// @Description create Random
// @Param	body		body 	models.Random	true		"body for Random content"
// @Success 201 {object} models.Random
// @Failure 403 body is empty
// @router / [post]
func (c *RandomController) Post() {
	var rsp randomResponse
	randomNum, err := c.Client.Random()
	if err != nil {
		rsp.Error = err.Error()
	} else {
		rsp.Error = "success"
		rsp.RandomNum = randomNum
	}

	b, _ := json.Marshal(rsp)
	c.Ctx.Output.Body(b)
}

// GetOne ...
// @Title GetOne
// @Description get Random by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Random
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RandomController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Random
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Random
// @Failure 403
// @router / [get]
func (c *RandomController) GetAll() {
	var rsp randomResponse

	rsp.Error = "success"
	rsp.RandomNum = c.RandomNum

	b, _ := json.Marshal(rsp)
	c.Ctx.Output.Body(b)
}

// Put ...
// @Title Put
// @Description update the Random
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Random	true		"body for Random content"
// @Success 200 {object} models.Random
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RandomController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Random
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RandomController) Delete() {

}
