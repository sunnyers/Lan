package components

import (
	"../lan"
	"fmt"
	"encoding/json"
)

type MobileApiControllerInterface interface {
	CheckLogin()
}

type MobileApiController struct {
	lan.Controller
	params map[string]interface{}
}

//执行控制器之前运行的操作
func (c *MobileApiController) BeforeController() {
	c.initCommonParams()
	if app, ok := c.AppController.(MobileApiControllerInterface); ok {
		app.CheckLogin()
	}
}

//初始化请求参数 merge(post,get)
func (c *MobileApiController) initCommonParams(){

}

//验证是否登录 暂未实现
func (c *MobileApiController) CheckLogin(){}

func (c *MobileApiController) ReturnSuccess() {
	c.ReturnJson(200, c.Data["data"], "")
}

// 返回json数据
func (c *MobileApiController) ReturnJson(code int64,data interface{},message string)  {
	returnData := make(map[string]interface{})
	returnData["code"] = code
	returnData["data"] = data
	returnData["message"] = message
	c.Ctx.W.Header().Set("Content-type","application/json; charset=utf-8")
	c.Ctx.W.Header().Set("Server","Tengine/2.0")
	if b, ok := json.Marshal(returnData); ok == nil {
		fmt.Fprintln(c.Ctx.W,string(b))
		return
	}
	fmt.Fprintln(c.Ctx.W,"{\"code\":78900,\"data\":[],\"message\":\"parsing json error\"}")
}

