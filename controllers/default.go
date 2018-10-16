package controllers

type MainController struct {
	MultiController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	//c.Ctx.Output.Download("/tmp/files/q")
}

func (this *MainController) Post() {
	this.SaveFiles("file_loader", "/tmp/files")
}

