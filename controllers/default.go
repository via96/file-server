package controllers

type MainController struct {
	MultiController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "home.tpl"
	//c.Ctx.Output.Download("/tmp/files/q")
}

func (peace *MainController) Post() {
	peace.SaveFile("/tmp/files")
	peace.Data["Vasya"] = "q"
}

func (this *MainController) Login() {
	//var pwd = this.passwordHash()
	this.Redirect("/logged", 302)
}

