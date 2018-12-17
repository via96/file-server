package controllers

import (
	"beego-fileServer/models/helpers"
	"file-server/models"
	"strings"
)

type LoginController struct {
	MultiController
}


func (this *LoginController) Get() {

	if this.isUserLogedIn() {
		curUser := this.getCurrentUser()
		this.Redirect("/user/" + curUser.Login, 302)
	} else {
		this.Redirect("/login", 302)
	}
}

func (this *LoginController) Post() {

}

func (this *LoginController) Login() {
	this.TplName = "login.tpl"
	if this.isUserLogedIn() {
		// редирект на страницу пользователя
	}

	if this.Ctx.Input.Method() == "GET" {
		return
	}

	login := this.GetString("login", "")
	pass := this.GetString("password", "")
	hashPass := this.passwordHash([]byte(pass))

	var o = helpers.GetORM()

	var user models.User
	err := o.QueryTable(new(models.User)).Filter("login", login).Filter("password", hashPass).One(&user)

	if err != nil {
		var sess = this.StartSession()
		defer sess.SessionRelease(this.Ctx.ResponseWriter)
		sess.Set("userId", user.Id)
		// редирект на страницу пользователя
	} else {
		this.Data["error"] = "Неверное имя пользователя или пароль"
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
		return
	}

	this.TplName = "register.tpl"

	login := this.GetString("login", "")
	pass := this.GetString("password", "")
	repass := this.GetString("repassword", "")

	o := helpers.GetORM()
	login = strings.ToLower(login)
	exist := o.QueryTable(new(models.User)).Filter("login", login).Exist()
	if exist {
		this.Data["error"] = "Пользователь с таким именем уже существует"
		return
	}

	if pass == repass {
		hashPass := this.passwordHash([]byte(pass))
		user := models.User{Login:login, Password:hashPass}

		if id, err := o.Insert(&user); err == nil {
			sess := this.StartSession()
			defer sess.SessionRelease(this.Ctx.ResponseWriter)
			sess.Set("userId", id)
			// редирект на страницу пользователя
		} else {
			this.Data["Error"] = err
		}

	}else {
		this.Data["Error"] = "Пароли не совпадают"
	}
}