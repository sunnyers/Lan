package main

import (
	_ "./routers"
	Lan "./lan"
	"os"
	"encoding/json"
	"fmt"
	"./components"
	"path/filepath"
	"log"
	"strings"
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
	RunTimes	string
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
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
		Lan.LogError("Init Wrong configuration: %s\n",err.Error())
		return
	}
	wd,err := os.Getwd()
	if err != nil{
		wd = getCurrentDirectory()
	}
	Lan.BackHost = conf.JavaApi.BackHost
	Lan.TimeOut = conf.JavaApi.TimeOut
	Lan.PathRoot = wd
	Lan.RunTimes = conf.RunTimes
	components.ValidTime = conf.ApiConfig.ValidTime
	components.CheckSign = conf.ApiConfig.CheckSign
	Lan.Run(conf.ServerAdd)
}