package controllers

import (
  "encoding/json"
	"github.com/astaxie/beego"
  "gopkg.in/mgo.v2/bson"
  "mealplans/models"
  "mealplans/dao"
  "golang.org/x/crypto/bcrypt"
)

var udao = dao.UserDAO{}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


type MainController struct {
	beego.Controller
}

type APIController struct {
  beego.Controller
}

func (this *MainController) Get() {
  this.Ctx.WriteString("The mealplan API is running.")
}


func (this *APIController) PopulateDB() {
  //TODO: create a MealPlan first
  password := "Apassword"
  passwordHash, _ := HashPassword(password)
  u := models.User {
    Id: bson.NewObjectId(),
    Name: "User1",
    Username: "user1",
    School: "Uni!",
    Password: passwordHash,
    //TODO: MealPlanId

  }
  err := udao.Insert(u)
  if err != nil {
    this.Ctx.WriteString("error")
    return
  }
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

/*
func (this *APIController) CreateUser() {
  name := this.Ctx.Input.Query("name")
  //school := this.Ctx.Input.Query("school")
  //username := this.Ctx.Input.Query("username")
  //password := this.Ctx.Input.Query("password")
  this.Ctx.WriteString(name)
}
*/
