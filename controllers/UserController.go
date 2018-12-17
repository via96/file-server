package controllers

import (
	_ "beego-fileServer/models/helpers"
	"file-server/models"
	_ "github.com/astaxie/beego/orm"
	"os"
)

type UserController struct {
	MultiController
}

func (this *UserController) Get() {
	if !this.isUserLogedIn() {
		this.Redirect("/", 302)
	}

	this.TplName = "home.tpl"
}

func (this *UserController) Post() {

}

func (this *UserController) Logout() {
	sess := this.StartSession()
	sess.Delete("userId")
	sess.SessionRelease(this.Ctx.ResponseWriter)
	this.Redirect("/", 302)
}

func (this *UserController) List() {
	var user = this.getCurrentUser()
	var o = this.getORM()
	this.TplName = "home.tpl"
	this.Data["CurrentUser"] = user

	var links []models.Link
	o.QueryTable(new(models.Link)).Filter("user_id", user.Id).OrderBy("upload_time").All(&links)

	this.Data["Files"] = links
	this.Data["Val"] = len(links)
}

func (this *UserController) Upload() {
	if !this.isUserLogedIn() {
		this.Redirect("/", 302)
	}

	this.TplName = "home.tpl"
	var user = this.getCurrentUser()

	if err := this.SaveFiles("file_loader", "/tmp/files"); err != nil {
		this.Data["Error"] = err
	} else {
		this.Redirect("/user/"+user.Login, 302)
	}
}


func (this *UserController) Download() {

	fileMarker := this.Ctx.Input.Param(":linkId")
	this.TplName = "home.tpl"
	this.Data["Website"] = fileMarker

	var o = this.getORM()
	var currentLink models.Link
	o.QueryTable(new(models.Link)).Filter("id", fileMarker).One(&currentLink)

	var targetFile models.File
	o.QueryTable(new(models.File)).Filter("id", currentLink.FileId).One(&targetFile)

	this.Data["Email"] = targetFile.Path
	this.Ctx.Output.Download(targetFile.Path, targetFile.Name)

	this.Redirect("/", 302)
}

func (this *UserController) Remove() {
	fileMarker := this.Ctx.Input.Param(":linkId")
	this.TplName = "home.tpl"
	var o = this.getORM()
	var currentLink models.Link
	o.QueryTable(new(models.Link)).Filter("id", fileMarker).One(&currentLink)

	o.Delete(&currentLink)

	var linkList []models.Link
	o.QueryTable(new(models.Link)).Filter("file_id", currentLink.FileId).All(&linkList)

	if len(linkList) == 0 {
		var currentFile models.File
		o.QueryTable(new(models.File)).Filter("id", currentLink.FileId).One(&currentFile)
		o.Delete(&currentFile)
		os.Remove(currentFile.Path)
	}
	this.Redirect("/", 302)
}