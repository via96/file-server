package models

import (
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id int64 `orm:"auto"`//`form:"-"`
	Login string //`form:"login,text,login:" valid:"MinSize(1);MaxSize(20)"`
	Password string //`form:"password,text,password:" valid:"MinSize(1);MaxSize(100)"`
}

type File struct {
	Id int64 `orm:"auto"`
	Name string //`form:"name,text,name:" valid:"MinSize(1);MaxSize(100)"`
	Hash string //`form:"hash,text,hash:" valid:"MinSize(1);MaxSize(100)"`
	Path string //`form:"path,text,path:" valid:"MinSize(1);MaxSize(200)"`
	UserId int64 //`form:"userId,integer,userId:"`
}

type Link struct {
	Id int64 `orm:"auto"`
	UserId int64
	FileId int64
	UserFileName string
	UploadTime string
}

func (u *User) TableName() string {
	return "User"
}

func (f * File) TableName() string {
	return "File"
}

func (uf * Link) TableName() string {
	return "Link"
}

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/test.db")

	orm.RegisterModel(new(User))
	orm.RegisterModel(new(File))
	orm.RegisterModel(new(Link))
}