package controllers

import (
	"github.com/astaxie/beego"
  "mealplans/dao"
)

var udao = dao.UserDAO{}

type MainController struct {
	beego.Controller
}

type APIController struct {
  beego.Controller
}

func (this *MainController) Get() {
  this.Ctx.WriteString("The mealplan API is running.")
}


func (this *APIController) GetAllUsers() {
}
