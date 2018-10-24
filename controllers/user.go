package controllers

import (
	"github.com/astaxie/beego"
  "encoding/json"
  "mealplans/models"
  "gopkg.in/mgo.v2/bson"
)

type UserAPIController struct {
  beego.Controller
}

func (this *UserAPIController) GetAllUsers() {
  users, err := udao.FindAll()
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  usersJson, err := json.Marshal(users)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  this.Ctx.WriteString(string(usersJson))
}

func (this *UserAPIController) CreateUser() {
  name := this.Ctx.Input.Query("name")
  school := this.Ctx.Input.Query("school")
  username := this.Ctx.Input.Query("new_username")
  passwordHash, _ := HashPassword(this.Ctx.Input.Query("new_password"))
  user := models.User {
    Id: bson.NewObjectId(),
    Name: name,
    Username: username,
    School: school,
    Password: passwordHash,
    IsAdmin: false,
    Token: GenerateUserToken(),
  }
  err := udao.Insert(user)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  userJson, err := json.Marshal(user)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  this.Ctx.WriteString(string(userJson))
}

