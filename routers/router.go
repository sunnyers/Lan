package routers

import (
	"../controller"
	"../controller/v1"
	"../lan"
)

func init(){
	lan.AutoRouter(&controller.DefaultController{})
	lan.NameSpaceRouter("/v1",
		&v1.DefaultController{},
		&v1.ApiController{},
	)
}