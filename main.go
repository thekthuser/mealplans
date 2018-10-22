package main

import (
	_ "mealplans/routers"
	"github.com/astaxie/beego"
  "mealplans/dao"
)

var udao = dao.UserDAO{}

func main() {
  udao.Server = "localhost"
  udao.Database = "mealplans_db"
  udao.Connect()
	beego.Run()
}

