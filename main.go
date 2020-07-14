package main

import (
	"github.com/astaxie/beego"
	"mvc/initial"
	_ "mvc/models"
	_ "mvc/routers"
)

func main() {
	//Db init
	initial.InitDb()
	beego.Run()

}
