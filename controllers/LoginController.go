package controllers

import (
	"file-server/models"
	"strings"
)

type LoginController struct {
	MultiController
}


func (this *LoginController) Get() {

	if this.isUserLogedIn() {
		curUser := this.getCurrentUser()
		this.Redirect("/user/"+curUser.Login, 302)
	}
	this.Redirect("/login", 302)
	//this.TplName = "login.tpl"
}

func (this *LoginController) Login() {
	this.TplName = "login.tpl"
	if this.isUserLogedIn() {
		curUser := this.getCurrentUser()
		this.Redirect("/user/"+curUser.Login, 302)
	}

	if this.Ctx.Input.Method() == "GET" {
		print("!!!! LOGIN GET !!!!")
		return
	}
	print("!!!! LOGIN POST !!!!")
	login := this.GetString("username_log", "")
	pass := this.GetString("password", "")
	hashPass := this.passwordHash([]byte(pass))

	var o = this.getORM()

	var user models.User
	err := o.QueryTable(new(models.User)).Filter("login", login).Filter("password", hashPass).One(&user)

	if err == nil {
		var sess = this.StartSession()
		defer sess.SessionRelease(this.Ctx.ResponseWriter)
		sess.Set("userId", user.Id)
		curUser := this.getCurrentUser()
		this.Redirect("/user/"+curUser.Login, 302)
	} else {
		this.Data["Error"] = "Неверное имя пользователя или пароль"
	}

}

func (this *LoginController) Logout() {
	sess := this.StartSession()
	sess.Delete("userId")
	sess.SessionRelease(this.Ctx.ResponseWriter)
	this.Redirect("/", 302)
}

func (this *LoginController) Register() {
	if this.isUserLogedIn() {
		this.Redirect("/", 302)
	}
	if this.Ctx.Input.Method() == "GET" {
		print("!!!! REGISTER GET !!!!")
		return
	}
	print("!!!! REGISTER POST !!!!")
	this.TplName = "login.tpl"

	login := this.GetString("username_reg", "")
	pass := this.GetString("password", "")
	repass := this.GetString("repassword", "")

	o := this.getORM()
	login = strings.ToLower(login)
	exist := o.QueryTable(new(models.User)).Filter("login", login).Exist()
	if exist {
		this.Data["Error"] = "Пользователь с таким именем уже существует"
		return
	}

	if pass == repass {
		hashPass := this.passwordHash([]byte(pass))
		user := models.User{Login:login, Password:hashPass}

		if id, err := o.Insert(&user); err == nil {
			sess := this.StartSession()
			defer sess.SessionRelease(this.Ctx.ResponseWriter)
			sess.Set("userId", id)
			this.Redirect("/", 302)
		} else {
			this.Data["Error"] = err
		}

	}else {
		this.Data["Error"] = "Пароли не совпадают"
	}
}