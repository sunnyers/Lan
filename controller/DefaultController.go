package controller

import (
	"../components"
	"fmt"
	"../common/dal"
)

type DefaultController struct {
	components.MobileApiController
	appIao dal.AppIao
}

func (c *DefaultController) Index(){
	c.Data["name"] = "this Index"
	fmt.Fprintln(c.Ctx.W,c.Data["name"])
}

func (c *DefaultController) Default() {
	params := make(map[string]interface{})
	params["appEnname"] = "wm_android"
	responseMessage := make(chan map[string]interface{})
	c.appIao.Request("getApp", params, responseMessage)
	response := <- responseMessage
	if response["code"].(int) == 200 {
		c.ReturnJson(200, response["data"], "")
	} else {
		c.ReturnJson(response["code"].(int), response["data"], response["message"].(string))
	}
}

func (c *DefaultController) GetList(){
	c.Data["name"] = "this is getList"
	fmt.Fprintln(c.Ctx.W,c.Data["name"])
}