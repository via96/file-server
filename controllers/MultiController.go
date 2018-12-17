package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"io"
	"os"
	"file-server/models"
)

type MultiController struct {
	beego.Controller
}

func (this *MultiController) SaveFiles(fromfile, dirName string) error {
	var files, err1 = this.GetFiles("file_loader")
	if err1 != nil {
		return err1
	}
	for _, header := range files  {
		file, _ := header.Open()
		f, err2 := os.OpenFile(dirName + "/" + header.Filename, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0777)
		if err2 != nil {
			return err2
		}
		defer f.Close()
		io.Copy(f, file)
	}
	return nil
}

func getORM() orm.Ormer  {
	o := orm.NewOrm()
	o.Using("default")
	return o
}

func (this *MultiController) isUserLogedIn() bool {
	sess := this.StartSession()
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	userId := sess.Get("userId")
	return userId != nil
}

func (this *MultiController) userExists(userId int) bool {
	sess := this.StartSession()
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	id := sess.Get("userId")
	return userId == id
}

func (this *MultiController) getCurrentUser() models.User {
	sess := this.StartSession()
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	userId := sess.Get("userId")
	o := orm.NewOrm()
	o.Using("default")
	var user models.User

	if err := o.QueryTable(new(models.User)).Filter("id", userId).One(&user); err != nil {
		this.Redirect("/login", 302)
	}
	return user
}

func (this *MultiController) passwordHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}