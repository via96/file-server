package main

import (
	_ "file-server/routers"
	"github.com/astaxie/beego"
	"os"
)

func main() {
	if _, err := os.Stat("/tmp/files"); os.IsNotExist(err) {
		os.Mkdir("/tmp/files", 0777)
	}
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

