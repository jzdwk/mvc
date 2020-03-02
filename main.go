package main

import (
	"github.com/astaxie/beego"
	"github.com/mvc/initial"
)

func main() {
	//Db init
	initial.InitDb()
	beego.Run()
}
