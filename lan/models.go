package lan

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"time"
	"encoding/json"
	"strconv"
	"github.com/muesli/cache2go"
	"errors"
	"sync"
)

var BackHost = ""
var TimeOut = 0
var client = &http.Client{Timeout: time.Second * time.Duration(TimeOut)}
var cache = cache2go.Cache("goCache")


var lock = new(sync.RWMutex)


var models *CModel

type myStruct struct {
	text     string
	moreData []byte
}

type ModelInterface interface {
	GetUri() string
	GetRules() [][]string
	GetCacheModel(name string) (bool,map[string]string)
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

func (c CModel) GetCacheModel(appIao ModelInterface,name string) (bool,map[string]string){
	return appIao.GetCacheModel(name)
}

func (c *CModel) Validate() bool {
	return true
}

func (c *CModel) Post(appIao ModelInterface,methodName string,params map[string]interface{},response chan map[string]interface{}) {
	result := make(map[string]interface{})
	parameter, _ := json.Marshal(params)
	requestUri := c.GetApiUrl(appIao) + methodName
	body := make([]byte,0)
	err := errors.New("ss")
	isCache := false
	ok,caches := c.GetCacheModel(appIao,methodName)
	if ok && caches["type"] == "memory"{
		isCache = true
		res, err := cache.Value(methodName)
		if err == nil {
			body = res.Data().([]byte)
		}
	}
	if len(body) == 0 {
		lock.Lock()
		if isCache && caches["type"] == "memory"{
			res, err := cache.Value(methodName)
			if err == nil {
				body = res.Data().([]byte)
			}
		}
		if len(body) == 0 {
			body, err = c.httpRequest(requestUri, parameter)
			result["code"] = 500
			result["message"] = "系统错误"
			if err != nil {
				fmt.Println("error is ", err.Error())
				response <- result
				return
			}
			if isCache {
				fmt.Println("set cache success ")
				if cacheTime,err := strconv.ParseInt(caches["time"],10,64); err == nil{
					cache.Add(methodName, time.Duration(cacheTime) * time.Second, body)
				}else{
					cache.Add(methodName, 120 * time.Second, body)
				}
			}
		}
		lock.Unlock()
	}
	js := json.NewDecoder(bytes.NewReader(body))
	js.UseNumber()
	errs := js.Decode(&result)
	if errs != nil{
		response <- result
		return
	}
	result["code"],_ = strconv.Atoi(string(result["code"].(json.Number)))
	if result["code"] == 500 || result["code"] == 501 || result["code"] == 502 {
		result["message"] = "您已退出登陆，请刷新页面后重新登录"
		response <- result
		return
	}
	response <- result
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
