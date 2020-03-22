package main

import (
	"github.com/astaxie/beego"
	"github.com/mvc/initial"
	_ "github.com/mvc/models"
	_ "github.com/mvc/routers"
)

func main() {
	//Db init
	initial.InitDb()
	beego.Run()

}
