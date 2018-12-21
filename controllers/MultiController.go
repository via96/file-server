package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"file-server/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io"
	"io/ioutil"
	"os"
	"time"
)

type MultiController struct {
	beego.Controller
}

func (this *MultiController) SaveFile(dirName string) error  {

	file, header, errGet := this.GetFile("file_loader")

	var user = this.getCurrentUser()
	var o = this.getORM()

	if errGet != nil {
		return errGet
	}

	if file != nil {

		file, errOpen := header.Open()

		if errOpen != nil {
			return errOpen
		}

		fileData, _ := ioutil.ReadAll(file)
		var hasherMd5 = md5.New()
		hasherMd5.Write(fileData)

		md5Hash := hex.EncodeToString(hasherMd5.Sum(nil))
		if dirName[len(dirName) - 1] == '/' {
			dirName = dirName[:len(dirName)-1]
		}
		targetPath := dirName + "/" + md5Hash

		file.Close()

		var fileItem models.File
		err := o.QueryTable(new(models.File)).Filter("hash", md5Hash).One(&fileItem)
		if err == orm.ErrNoRows {
			if _, err := os.Stat(dirName); os.IsNotExist(err) {
				os.Mkdir(dirName, 0777)
			}
			errSave := this.SaveToFile("file_loader", "/tmp/files/" + md5Hash)
			if errSave != nil {
				return errSave
			}

			fileToAdd := models.File{Hash:md5Hash, Path:targetPath, UserId:user.Id, Name:header.Filename}
			if id, err := o.Insert(&fileToAdd); err == nil{
				fileToAdd.Id = id
				this.linkFileToUser(&user, &fileToAdd, header.Filename)
				return nil
			}
		}

		if err == orm.ErrMultiRows {
			return nil
		}
		this.linkFileToUser(&user, &fileItem, header.Filename)
	}
	return nil
}

func (this *MultiController) SaveFiles_old(fromfile, dirName string) error {
	var files, err1 = this.GetFiles("file_loader")
	if err1 != nil {
		return err1
	}

	var user = this.getCurrentUser()
	var o = this.getORM()

	for _, header := range files {
		file, _ := header.Open()

		fileData, _ := ioutil.ReadAll(file)
		var hasherMd5 = md5.New()
		hasherMd5.Write(fileData)

		md5Hash := hex.EncodeToString(hasherMd5.Sum(nil))

		var fileItem models.File
		err := o.QueryTable(new(models.File)).Filter("hash", md5Hash).One(&fileItem)
		if err == orm.ErrNoRows {
			targetPath := dirName + "/" + md5Hash
			f, err2 := os.OpenFile(targetPath, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0777)
			if err2 != nil {
				return err2
			}
			defer f.Close()
			io.Copy(f, file)
			fileToAdd := models.File{Hash:md5Hash, Path:targetPath, UserId:user.Id, Name:header.Filename}
			if id, err := o.Insert(&fileToAdd); err == nil{
				fileToAdd.Id = id
				this.linkFileToUser(&user, &fileToAdd, header.Filename)
				return nil
			}
		}
		if err == orm.ErrMultiRows {
			return nil
		}
		this.linkFileToUser(&user, &fileItem, header.Filename)
	}
	return nil
}

func (this *MultiController) linkFileToUser(user *models.User, file *models.File, filename string) {
	var o = this.getORM()

	var link = models.Link{UserId:user.Id, FileId:file.Id, UserFileName:filename, UploadTime: time.Now().Format(time.RFC3339)}
	o.Insert(&link)
}

func (this *MultiController) getORM() orm.Ormer  {
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
	o := this.getORM()
	var user models.User

	if err := o.QueryTable(new(models.User)).Filter("id", userId).One(&user); err != nil {
		this.Redirect("/login", 302)
	}
	return user
}

func (this *MultiController) passwordHash(pwd []byte) string {
	var hasherMd5 = md5.New()
	hasherMd5.Write(pwd)

	hash := hex.EncodeToString(hasherMd5.Sum(nil))

	return string(hash)
}

func (this *MultiController) getFileList()  {
	sess := this.StartSession()
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	userId := sess.Get("userId")
	o := orm.NewOrm()
	o.Using("default")
	var params []orm.ParamsList
	if num, err:= o.QueryTable(new(models.Link)).Filter("id", userId).ValuesList(&params);  err == nil && num > 0 {
		for val := range params {
			print(val)
		}
	}
}