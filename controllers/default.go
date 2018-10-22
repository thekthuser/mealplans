package controllers

import (
  "encoding/json"
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
  users, err := udao.FindAll()
  if err != nil {
    this.Ctx.WriteString("error")
    return
  }
  /*
  var output string
  for _, user := range users {
    line, _ := json.Marshal(user)
    output = output + "<br />" + string(line)
  }
  this.Ctx.WriteString(output)
  */
  usersJson, err := json.Marshal(users)
  if err != nil {
    this.Ctx.WriteString("error")
    return
  }
  this.Ctx.WriteString(string(usersJson))
}
