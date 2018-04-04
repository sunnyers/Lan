package lan

import (
	"sync"
	"net/http"
	"reflect"
	"strings"
)

const defController  = "default"
const defMethod  = "Index"

var routers map[string]reflect.Type = make(map[string]reflect.Type)
var namespace map[string]bool = make(map[string]bool)

type handler struct {
	pool         sync.Pool
}

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (ctx *Context) Config(w http.ResponseWriter, r *http.Request){
	ctx.R = r
	ctx.W = w
}

func NewHandler() *handler {
	h := &handler{}
	h.pool.New = func() interface{} {
		return &Context{}
	}
	return h
}


func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx := h.pool.Get().(*Context)
	defer h.pool.Put(ctx)
	ctx.Config(w, r)

	controllerName, methodName := h.findControllerInfo(w,r)
	controllerT, ok := routers[controllerName]
	if !ok {
		http.NotFound(w, r)
		return
	}

	refV := reflect.New(controllerT)
	method := refV.MethodByName(methodName)
	if !method.IsValid() {
		http.NotFound(w, r)
		return
	}
	controller := refV.Interface().(ControllerInterface)
	controller.Init(ctx,controllerName,methodName,refV.Interface())
	controller.BeforeController()
	controller.BeforeAction()
	method.Call(nil)
}

func (h *handler) findControllerInfo(w http.ResponseWriter,r *http.Request) (string, string) {
	path := r.URL.Path
	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}
	pathInfo := strings.Split(path, "/")
	isName := false
	controllerName := defController
	methodName := defMethod
	num := len(pathInfo)
	if num > 1 {
		controllerName = strings.ToLower(pathInfo[1])
		if _,ok := namespace["/"+ controllerName]; ok{
			if num > 2 {
				controllerName = "/" + controllerName + "/" + strings.ToLower(pathInfo[2])
				isName = true
			}else{
				controllerName = "/" + controllerName + "/" + defController
				isName = true
			}
		}
	}
	if isName && num > 3 {
		methodName = strings.Title(pathInfo[3])
	}
	if !isName && num > 2 {
		methodName = strings.Title(pathInfo[2])
	}
	return controllerName, methodName
}

func router(pattern string, t reflect.Type) {
	routers[strings.ToLower(pattern)] = t
}

func Router(pattern string, app ControllerInterface) {
	refV := reflect.ValueOf(app)
	refT := reflect.Indirect(refV).Type()
	router(pattern, refT)
}

func AutoRouter(app ControllerInterface) {
	refV := reflect.ValueOf(app)
	refT := reflect.Indirect(refV).Type()
	refName := strings.TrimSuffix(strings.ToLower(refT.Name()), "controller")
	router(refName, refT)
}

func NameSpaceRouter(pattern string,appList ...ControllerInterface){
	namespace[pattern] = true
	for _,app := range appList{
		refV := reflect.ValueOf(app)
		refT := reflect.Indirect(refV).Type()
		refName := strings.TrimSuffix(strings.ToLower(refT.Name()), "controller")
		router(pattern + "/" + refName, refT)
	}
}