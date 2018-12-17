package models

import (
	"beego-fileServer/models"
	"file-server/controllers"
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int64 `form:"-"`
	Login string `form:"login,text,login:" valid:"MinSize(5);MaxSize(20)"`
	Password string `form:"login,text,login:" valid:"MinSize(1);MaxSize(100)"`
}

type File struct {
	Id int64 `form:"-"`
	Name string `form:"name,text,name:" valid:"MinSize(1);MaxSize(100)"`
	Hash string `form:"hash,text,hash:" valid:"MinSize(1);MaxSize(100)"`
	Path string `form:"path,text,path:" valid:"MinSize(1);MaxSize(200)"`
	UserId int64 `form:"userId,integer,userId:"`
}

type Link struct {
	UserId int64
	FileId int64
	UserFileName string
	UploadTime string
}

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/file_server.db")

	orm.RegisterModel(new(User))
	orm.RegisterModel(new(File))
}

func addFile(name string, path string)  {
	o := orm.NewOrm()
	o.Using("default")

}

func getCurrentUser(this *controllers.MultiController) models.User {
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

func getFileList(this *controllers.MultiController)  {
	sess := this.StartSession()
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	userId := sess.Get("userId")
	o := orm.NewOrm()
	o.Using("default")
	var params []orm.ParamsList
	var fileList []File
	if num, err:= o.QueryTable(new(models.File)).Filter("id", userId).ValuesList(&params);  err == nil && num > 0 {
		for val := range params {
			print(val)
		}
	}
}