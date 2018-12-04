package controllers

type UserController struct {
	MultiController
}

func (c *UserController) Get() {
	c.TplName = "logged.tpl"
	//c.Ctx.Output.Download("/tmp/files/q")
}

func (peace *UserController) Post() {

}