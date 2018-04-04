package lan

type Controller struct {
	Ctx *Context
	Data map[string]interface{}
	ControllerName string
	MethodName string
	AppController interface{}
}

type ControllerInterface interface {
	Init(ctx *Context,controllerName string,methodName string,app interface{})
	BeforeController()
	BeforeAction()
}

func (c *Controller) Init(ctx *Context,controllerName string,method string,app interface{}){
	c.Ctx = ctx
	c.Data = make(map[string]interface{})
	c.ControllerName = controllerName
	c.MethodName = method
	c.AppController = app
}

func (c *Controller) GetParams(name string,defaultValue string) string {
	c.Ctx.R.ParseForm()
	get := make(map[string]interface{})
	if len(c.Ctx.R.Form) > 0 {
		for k, v := range c.Ctx.R.Form {
			get[k] = v[0]
		}
	}
	post := make(map[string]interface{})
	if len(c.Ctx.R.PostForm) > 0 {
		for k, v := range c.Ctx.R.PostForm {
			post[k] = v[0]
		}
	}
	params := MapMerge(get,post)
	if value,ok := params[name]; ok && value != "" {
		return value.(string)
	}else{
		return defaultValue
	}
}

func (c *Controller) GetQuery(name string,defaultValue string) string {
	c.Ctx.R.ParseForm()
	get := c.Ctx.R.Form.Get(name)
	if get == "" {
		return defaultValue
	}
	return get
}

func (c *Controller) GetPost(name string,defaultValue string) string {
	get := c.Ctx.R.PostForm.Get(name)
	if get == "" {
		return defaultValue
	}
	return get
}

func (c *Controller) BeforeController(){}
func (c *Controller) BeforeAction(){}