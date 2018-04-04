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
	sp := c.GetParams("aa","bb")
	if sp == "100" {
		fmt.Println("int ",sp)
	}else{
		fmt.Println("interface ",sp)
	}
	code,rsp := c.appIao.Request("getApp", params)
	if code == 200 {
		c.ReturnJson(200, rsp["data"], "")
	} else {
		c.ReturnJson(code, rsp["data"], rsp["message"].(string))
	}
}

func (c *DefaultController) GetList(){
	c.Data["name"] = "this is getList"
	fmt.Fprintln(c.Ctx.W,c.Data["name"])
}