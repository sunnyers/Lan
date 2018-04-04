package v1

import (
	"../../components"
	"fmt"
)

type DefaultController struct {
	components.MobileApiController
}

func (c *DefaultController) Index(){
	c.Data["name"] = "this v1 Index"
	fmt.Fprintln(c.Ctx.W,c.Data["name"])
}