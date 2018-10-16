package main

import (
	_ "file-server/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

