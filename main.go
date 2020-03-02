package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "imooc/routers"
)

func main() {
	//orm.RegisterDataBase("default","mysql","root:123456@tcp(localhost:3306)/imooc?charset=utf8")
	beego.Run()
}
