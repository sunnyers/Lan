package dal

import (
	"../../lan"
	"sync"
)

var appIao *AppIao
var once sync.Once

type AppIao struct {
	lan.CModel
}

//单例模式
func NewAppIao() *AppIao{
	once.Do(func(){
		appIao = &AppIao{}
	})
	return appIao
}

func (c *AppIao) GetUri() string{
	return "app/"
}

// rules验证 暂未实现
func (c *AppIao) GetRules() [][]string{
	m := [][]string{
		{"objectId","required"},
		{"aaa","vvv"},
	}
	return m
}

func (c *AppIao) Request(methodName string,params map[string]interface{},response chan map[string]interface{}) {
	go c.Post(c,methodName, params,response)
}