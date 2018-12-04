package models

import "github.com/astaxie/beego/orm"

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

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/file_server.db")

	orm.RegisterModel(new(User))
	orm.RegisterModel(new(File))
}