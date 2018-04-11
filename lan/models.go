package lan

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"time"
	"encoding/json"
	"strconv"
)

var BackHost = ""
var TimeOut = 0
var client = &http.Client{Timeout: time.Second * 2}

var models *CModel

type ModelInterface interface {
	GetUri() string
	GetRules() [][]string
}

type CModel struct {
	Attributes map[string]interface{}
}

func (c *CModel) GetRules(appIao ModelInterface) map[int]interface{} {
	return make(map[int]interface{})
}

func (c *CModel) GetApiUrl(appIao ModelInterface) string {
	return BackHost + appIao.GetUri()
}

func (c *CModel) Validate() bool {
	return true
}

func (c *CModel) Post(appIao ModelInterface,methodName string,params map[string]interface{}) (int,map[string]interface{}) {
	result := make(map[string]interface{})
	parameter, _ := json.Marshal(params)
	requestUri := c.GetApiUrl(appIao) + methodName
	body,err := c.httpRequest(requestUri, parameter)
	code := 500
	result["message"] = "系统错误"
	if err != nil{
		fmt.Println("error is ",err.Error())
		return code,result
	}
	js := json.NewDecoder(bytes.NewReader(body))
	js.UseNumber()
	errs := js.Decode(&result)
	if errs != nil{
		return code,result
	}
	code,_ = strconv.Atoi(string(result["code"].(json.Number)))
	if code == 200 || code == 501 || code == 502 {
		result["message"] = "您已退出登陆，请刷新页面后重新登录"
		return code,result
	}
	delete(result,"code")
	return code,result
}

func (c *CModel) httpRequest(url string,parameter []byte) ([]byte, error) {
	body := make([]byte, 0)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(parameter))
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Connection", "close")
	response, err := client.Do(request)
	if err != nil {
		if nil != response {
			response.Body.Close()
		}
		fmt.Sprintln("httpReauest %s do err : %s - %d \n", url, err.Error())
		return nil, err
	}
	if response.StatusCode == 200 {
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			response.Body.Close()
			fmt.Sprintln("httpReauest %s ReadAll err : %s - %d\n", url, err.Error())
			return nil, err
		}
		fmt.Sprintln("recv http data reply: %s--%s--%s\n", url, string(body), string(parameter))
	} else {
		fmt.Sprintln("recv http response StatusCode !=200: %d-%s\n", response.StatusCode, url)
	}
	response.Body.Close()
	return body, nil
}
