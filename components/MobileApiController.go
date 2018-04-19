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
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"bytes"
	"sort"
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
		if CheckSign == true {
			return c.checkSign()
		}else {
			return true
		}
	}
	return false
}

func (c *MobileApiController) checkSign() bool{
	if c.checkAppName() {
		sign := c.params["sign"]
		if  c.createSign(c.params) == sign {
			return true
		}
		c.ReturnError(config.COMMUNAL_PARAM_SIGN_INVALID,"")
	}
	return false
}

func (c *MobileApiController) createSign(params map[string]interface{}) string {
	delete(params,"sign")
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)  //升序
	//sort.Sort(sort.Reverse(sort.StringSlice(keys)))  //倒序
	var buf bytes.Buffer
	for _,k := range keys {
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(params[k].(string))
		buf.WriteString("&")
	}
	signStr := buf.String()
	signStr = signStr[0:len(signStr)-1]
	s := url.QueryEscape(signStr)
	m := md5.New()
	m.Write([]byte(s))
	md5Value := hex.EncodeToString(m.Sum(nil))
	return md5Value
}

func (c *MobileApiController) checkAppName() bool {
	params := make(map[string]interface{})
	params["appEnname"] = c.params["appkey"]
	var appIao Dal.AppIao
	responseChan := make(chan map[string]interface{})
	appIao.Request("getApp", params, responseChan)
	response := <-responseChan
	if response["code"].(int) == 200 {
		data := response["data"].(map[string]interface{})
		if response["data"] == "" || data["appStatus"] == 0 {
			c.ReturnError(config.COMMUNAL_PARAM_APPKEY_INVALID, "")
		}
		c.params["appsecret"] = data["appSecret"]
		return true
	} else {
		c.ReturnJson(response["code"].(int), make([]interface{}, 0), response["message"].(string))
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

