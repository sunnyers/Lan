package lan

import (
	"log"
	"os"
	"time"
	"errors"
)

var PathRoot string = ""
var RunTimes string = ""
var (
	Trace *log.Logger
	Info *log.Logger
	Error *log.Logger
	EFD *os.File
	IFD *os.File
	TFD *os.File
)

// 判断文件夹是否存在
func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func LogInfo(format string, v ...interface{}){
	filePath := PathRoot + RunTimes + time.Unix(time.Now().Unix(), 0).Format("2006/01/02")
	filename := filePath + "/info.log"
	err := errors.New("")
	if checkFileIsExist(filename) {
		IFD, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("open file error !", filename,err.Error())
		}
	}else{
		if !pathExists(filePath){
			err := os.MkdirAll(filePath, os.ModePerm)
			log.Fatalln("mkdir error ",err.Error())
		}
		IFD, err = os.Create(filename) //创建文件
		if err != nil {
			log.Fatalln("create file error !", filename,err.Error())
		}
	}
	Info = log.New(IFD, "[Info] ", log.LstdFlags)
	Info.Printf(format, v...)
}

func LogError(format string, v ...interface{}){
	filePath := PathRoot + RunTimes + time.Unix(time.Now().Unix(), 0).Format("2006/01/02")
	filename := filePath + "/error.log"
	err := errors.New("")
	if checkFileIsExist(filename) {
		EFD, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("open file error !", filename,err.Error())
		}
	}else{
		if !pathExists(filePath){
			err := os.MkdirAll(filePath, os.ModePerm)
			log.Fatalln("mkdir error ",err.Error())
		}
		EFD, err = os.Create(filename) //创建文件
		if err != nil {
			log.Fatalln("create file error !", filename,err.Error())
		}
	}
	Error = log.New(EFD, "[Error] ", log.LstdFlags)
	Error.Printf(format, v...)
}


func LogTrace(format string, v ...interface{}){
	filePath := PathRoot + RunTimes + time.Unix(time.Now().Unix(), 0).Format("2006/01/02")
	filename := filePath + "/error.log"
	err := errors.New("")
	if checkFileIsExist(filename) {
		TFD, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("open file error !", filename,err.Error())
		}
	}else{
		if !pathExists(filePath){
			err := os.MkdirAll(filePath, os.ModePerm)
			log.Fatalln("mkdir error ",err.Error())
		}
		TFD, err = os.Create(filename) //创建文件
		if err != nil {
			log.Fatalln("create file error !", filename,err.Error())
		}
	}
	Trace = log.New(TFD, "[Trace] ", log.LstdFlags)
	Trace.Printf(format, v...)
}
