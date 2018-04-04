package main

import (
	_ "./routers"
	Lan "./lan"
	"os"
	"encoding/json"
	"fmt"
)

type javaApi struct {
	BackHost	string
	TimeOut 	int
}

type configuration struct {
	ServerAdd	string
	IsCheck		bool
	JavaApi		javaApi
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
	Lan.Run(conf.ServerAdd)
}