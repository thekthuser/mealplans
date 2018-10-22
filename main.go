package main

import (
	_ "mealplans/routers"
	"github.com/astaxie/beego"
  "mealplans/dao"
)

var udao = dao.UserDAO{}
var pdao = dao.PlanDAO{}

func main() {
  udao.Server = "localhost"
  udao.Database = "mealplans_db"
  udao.Connect()
  pdao.Server = "localhost"
  pdao.Database = "mealplans_db"
  pdao.Connect()
	beego.Run()
}

