package components

import (
	"../lan"
	"fmt"
	"encoding/json"
	"strings"
	"../config"
	"strconv"
	"time"
	Dal "../common/dal"
)

var ValidTime int64 = 0
var CheckSign bool = true

type MobileApiControllerInterface interface {
	CheckLogin()
}

type MobileApiController struct {
	lan.Controller
	params map[string]interface{}
	clientVersion string //客户端版本
	clientType string //客户端类型 ios android
	channelType string //客户端渠道 yx wangmiao
	streamType string //视频流类型 rtmp flv
}

//执行控制器之前运行的操作
func (c *MobileApiController) BeforeController() bool {
	c.initCommonParams()
	if c.checkCommonParams() {
		return c.checkSign()
	}
	return false
}

func (c *MobileApiController) checkSign() bool{
	return c.checkAppName()


	return true
}

func (c *MobileApiController) checkAppName() bool {
	params := make(map[string]interface{})
	params["appEnname"] = c.params["appkey"]
	var appIao Dal.AppIao
	code,rsp := appIao.Request("getApp", params)
	if code == 200 {
		data := rsp["data"].(map[string]interface{})
		if rsp["data"] == "" || data["appStatus"] == 0 {
			c.ReturnError(config.COMMUNAL_PARAM_APPKEY_INVALID, "")
		}
		c.params["ppsecret"] = data["ppSecret"]
		return true
	} else {
		c.ReturnJson(code, make([]interface{},0), rsp["message"].(string))
		return false
	}
}

func (c *MobileApiController) checkCommonParams() bool{
	if _,ok := lan.Isset(c.params,"appkey"); !ok {
		c.ReturnError(config.COMMUNAL_PARAM_APPKEY_INVALID,"")
		return false
	}
	if timestamp,ok := lan.Isset(c.params,"timestamp"); ok{
		t := time.Now()
		ts,_ := strconv.ParseInt(timestamp,10,64)
		diff := ts - t.Unix()
		if diff < 0 {
			diff = -diff
		}
		fmt.Println(ts,t.Unix(),diff)
		if diff > ValidTime {
			c.ReturnError(config.COMMUNAL_PARAM_TIMESTAMP_INVALID,"")
			return false
		}
	}else{
		c.ReturnError(config.COMMUNAL_PARAM_TIMESTAMP_INVALID,"")
		return false
	}
	if _,ok := lan.Isset(c.params,"session_id"); !ok{
		c.ReturnError(config.COMMUNAL_PARAM_SESSIONID_INVALID,"")
		return false
	}
	if _,ok := lan.Isset(c.params,"sign"); !ok{
		c.ReturnError(config.COMMUNAL_PARAM_SIGN_INVALID,"")
		return false
	}
	return true
}

//初始化请求参数 merge(post,get)
func (c *MobileApiController) initCommonParams() {
	c.params = c.GetMergeParams()
	if v,ok := lan.Isset(c.params,"v"); ok {
		f := strings.Split(v,"_")
		if len(f) > 2 {
			c.channelType = f[0]
			c.clientType = f[1]
			c.clientVersion = f[2]
		}
		delete(c.params,"v")
	}
}

//验证是否登录 暂未实现
func (c *MobileApiController) CheckLogin(){}

func (c *MobileApiController) ReturnSuccess() {
	c.ReturnJson(200, c.Data["data"], "")
}
func (c *MobileApiController) ReturnError(code int,message string)  {
	if message != ""{
		c.ReturnJson(code, make([]int,0), message)
		return
	}
	c.ReturnJson(code, make([]int,0), config.GetDesc(code))
}

// 返回json数据
func (c *MobileApiController) ReturnJson(code int,data interface{},message string)  {
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

