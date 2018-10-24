package controllers

import (
  "math/rand"
  "regexp"
  "time"
	"github.com/astaxie/beego"
  "gopkg.in/mgo.v2/bson"
  "mealplans/models"
  "mealplans/dao"
  "golang.org/x/crypto/bcrypt"
  "github.com/astaxie/beego/context"
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

func EmptyOrIsDate(date string) bool {
  if date == "" {
    return true
  }
  //check date is in MM/DD/YYYY format
  match, _ := regexp.MatchString("^[0-9]{2}/[0-9]{2}/[0-9]{4}$", date)
  return match
}

func GenerateUserToken() string {
  rand.Seed(time.Now().UTC().UnixNano())
  const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
  b := make([]byte, 32)
    for i := range b {
        b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
    }
    return string(b)
}

var LoginFilter = func(ctx *context.Context) {
  username := ctx.Input.Query("username")
  password := ctx.Input.Query("password")
  user, err := udao.FindByUsername(username)
  if err != nil {
    ctx.ResponseWriter.WriteHeader(401)
  }
  if !CheckPasswordHash(password, user.Password) {
    ctx.ResponseWriter.WriteHeader(401)
  }
}

var TokenFilter = func(ctx *context.Context) {
  user_id := ctx.Input.Param(":user_id")
  token := ctx.Input.Param(":token")
  if !bson.IsObjectIdHex(user_id) {
    ctx.ResponseWriter.WriteHeader(400)
    return
  }
  user, err := udao.FindById(user_id)
  if err != nil {
    ctx.ResponseWriter.WriteHeader(401)
  }
  if user.Token != token {
    ctx.ResponseWriter.WriteHeader(401)
  }
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
    this.Ctx.ResponseWriter.WriteHeader(500)
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
    MealPlanId: p.Id.Hex(),
    IsAdmin: true,
    Token: GenerateUserToken(),
  }
  err = udao.Insert(u)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  this.Ctx.WriteString("Database populated.")
}
