package v1

import (
	"../../components"
	"fmt"
)

type ApiController struct {
	components.MobileApiController
}

func (c *ApiController) Index(){
	c.Data["name"] = "this api Index"
	fmt.Fprintln(c.Ctx.W,c.Data["name"])
}
