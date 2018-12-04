package controllers

import (
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"io"
	"os"
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

func (this *MultiController) passwordHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}