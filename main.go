package main

import (
	_ "./routers"
	Lan "./lan"
	"os"
	"encoding/json"
	"fmt"
	"./components"
)

type javaApi struct {
	BackHost	string
	TimeOut 	int
}

type apiConfig struct {
	ValidTime	int64
	CheckSign	bool
}


type configuration struct {
	ServerAdd	string
	IsCheck		bool
	JavaApi		javaApi
	ApiConfig	apiConfig
}
func main(){
	file, err := os.Open("./config/config.json")
	if err != nil {
		fmt.Println("sad")
		return
	}
	decoder := json.NewDecoder(file)
	conf := configuration{}
	err = decoder.Decode(&conf)
	file.Close()
	if err != nil {
		fmt.Sprintln("Wrong configuration: %s\n", err.Error())
		return
	}
	Lan.BackHost = conf.JavaApi.BackHost
	Lan.TimeOut = conf.JavaApi.TimeOut
	components.ValidTime = conf.ApiConfig.ValidTime
	components.CheckSign = conf.ApiConfig.CheckSign
	Lan.Run(conf.ServerAdd)
}