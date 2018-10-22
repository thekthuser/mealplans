package controllers

import (
  "regexp"
  "encoding/json"
	"github.com/astaxie/beego"
  "gopkg.in/mgo.v2/bson"
  "mealplans/models"
  "mealplans/dao"
  "golang.org/x/crypto/bcrypt"
)

var udao = dao.UserDAO{}
var pdao = dao.PlanDAO{}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsDate(date string) bool {
  //check date is in MM/DD/YYYY format
  match, _ := regexp.MatchString("[0-9]{2}/[0-9]{2}/[0-9]{4}", date)
  return match
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
  p := models.Plan {
    Id: bson.NewObjectId(),
    Name: "Plan1",
    Cost: 25,
    Market: "new_york",
    Semester1Start: "01/01/2018",
    Semester1End: "02/01/2018",
    Semester2Start: "02/02/2018",
    Semester2End: "03/01/2018",
    Semester3Start: "03/02/2018",
    Semester3End: "04/01/2018",
    MarketingText1: "MarketingText1!",
    MarketingText2: "MarketingText2!",
    MarketingText3: "MarketingText3!",
  }
  err := pdao.Insert(p)
  if err != nil {
    this.Ctx.WriteString("error")
    return
  }
  password := "Apassword"
  passwordHash, _ := HashPassword(password)
  u := models.User {
    Id: bson.NewObjectId(),
    Name: "User1",
    Username: "user1",
    School: "Uni!",
    Password: passwordHash,
    MealPlanId: p.Id,
    IsAdmin: true,

  }
  err = udao.Insert(u)
  if err != nil {
    this.Ctx.WriteString("error")
    return
  }
  this.Ctx.WriteString("Database populated.")
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

func (this *APIController) GetAllPlans() {
  plans, err := pdao.FindAll()
  if err != nil {
    this.Ctx.WriteString("error")
    return
  }
  plansJson, err := json.Marshal(plans)
  if err != nil {
    this.Ctx.WriteString("error")
    return
  }
  this.Ctx.WriteString(string(plansJson))
}

func (this *APIController) GetAllPlansInMarket() {
  market := this.Ctx.Input.Param(":market")
  plans, err := pdao.FindByMarket(market)
  if err != nil {
    this.Ctx.WriteString("error")
    return
  }
  plansJson, err := json.Marshal(plans)
  if err != nil {
    this.Ctx.WriteString("error")
    return
  }
  this.Ctx.WriteString(string(plansJson))
}

func (this *APIController) GetPlan() {
  id := this.Ctx.Input.Param(":id")
  plan, err := pdao.FindById(id)
  if err != nil {
    this.Ctx.WriteString("error")
    return
  }
  planJson, err := json.Marshal(plan)
  if err != nil {
    this.Ctx.WriteString("error")
    return
  }
  this.Ctx.WriteString(string(planJson))
}
