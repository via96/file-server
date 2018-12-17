package controllers

import (
	"beego-fileServer/models/helpers"
	"crypto"
	"crypto/md5"
	"encoding/hex"
	"file-server/models"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"mime/multipart"
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

func (this *UserController) List() {
	var user = helpers.GetCurrentUser(&this.Controller)
	var o = helpers.GetORM()
	this.TplName = "home.tpl"
	//this.Data["items"] = user.Files

	//var files []*models.File
	//o.QueryTable(new(models.UserFile)).Filter("user_id", user.Id).OrderBy("upload_time").All(&files)
	//
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(
		"users_files.user_id",
		"users_files.file_id",
		"users_files.user_file_name",
		"users_files.upload_time",).
		From("users_files").
		InnerJoin("files").On("users_files.file_id = files.id").
		Where("users_files.user_id = ?").
		OrderBy("users_files.upload_time")

	var userFiles []models.Link

	sql := qb.String()
	o.Raw(sql, user.Id).QueryRows(&userFiles)

	this.Data["Files"] = userFiles
	this.Data["Val"] = len(userFiles)
}

func (this *UserController) Upload() {
	helpers.SetLayoutFor(&this.Controller)
	//var o = helpers.GetORM()
	var user = helpers.GetCurrentUser(&this.Controller)
	this.TplName = "forms/upload.tpl"

	if this.Ctx.Request.Method == "GET" {
		return
	}
	var files, err = this.GetFiles("the_file")

	if err != nil {
		return
	}
	for _, header:= range files {
		processFile(&user, header)
	}

	this.Redirect("/user/"+user.Login, 302)
}


func (this *UserController) Download() {

	fileMarker := this.Ctx.Input.Param(":name")
	this.TplName = "index.tpl"
	this.Data["Website"] = fileMarker

	var o = helpers.GetORM()
	var user = helpers.GetCurrentUser(&this.Controller)
	//this.Data["items"] = user.Files

	//var files []*models.File
	//o.QueryTable(new(models.UserFile)).Filter("user_id", user.Id).OrderBy("upload_time").All(&files)
	//
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(
		"users_files.user_id",
		"users_files.file_id",
		"users_files.user_file_name",
		"users_files.upload_time").
		From("users_files").
		InnerJoin("files").On("users_files.file_id = files.id").
		Where("users_files.user_id = ? AND users_files.upload_time = ?")

	var userFiles []models.Link
	sql := qb.String()
	o.Raw(sql, user.Id, fileMarker).QueryRows(&userFiles)
	this.Data["Website"] = len(userFiles)
	if len(userFiles) > 0 {
		this.Data["Email"] = userFiles[0].Stored
		this.Ctx.Output.Download(userFiles[0].Stored)
	}

	this.Redirect("/", 302)
}

func processFile(user *models.User, header *multipart.FileHeader) {
	file, errOpen := header.Open()
	if errOpen != nil {
		return
	}

	newFileData, _ := ioutil.ReadAll(file)
	var o = getORM()

	var hasherMd5 = md5.New()
	hasherMd5.Write(newFileData)

	md5Hash := hex.EncodeToString(hasherMd5.Sum(nil))

}