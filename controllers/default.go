package controllers

import (
  "math/rand"
  "regexp"
  "strconv"
  "encoding/json"
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
    MealPlanId: p.Id,
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

func (this *APIController) GetAllUsers() {
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
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  plansJson, err := json.Marshal(plans)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  this.Ctx.WriteString(string(plansJson))
}

func (this *APIController) GetAllPlansInMarket() {
  market := this.Ctx.Input.Param(":market")
  plans, err := pdao.FindByMarket(market)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  plansJson, err := json.Marshal(plans)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  this.Ctx.WriteString(string(plansJson))
}

func (this *APIController) GetPlan() {
  id := this.Ctx.Input.Param(":plan_id")
  plan, err := pdao.FindById(id)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  planJson, err := json.Marshal(plan)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  this.Ctx.WriteString(string(planJson))
}

func (this *APIController) CreatePlan() {
  name := this.Ctx.Input.Query("name")
  cost, err := strconv.Atoi(this.Ctx.Input.Query("cost"))
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(400)
    return
  }
  market := this.Ctx.Input.Query("market")
  semester1start := this.Ctx.Input.Query("semester1start")
  semester1end := this.Ctx.Input.Query("semester1end")
  semester2start := this.Ctx.Input.Query("semester2start")
  semester2end := this.Ctx.Input.Query("semester2end")
  semester3start := this.Ctx.Input.Query("semester3start")
  semester3end := this.Ctx.Input.Query("semester3end")
  marketingtext1 := this.Ctx.Input.Query("marketingtext1")
  marketingtext2 := this.Ctx.Input.Query("marketingtext2")
  marketingtext3 := this.Ctx.Input.Query("marketingtext3")
  if (!EmptyOrIsDate(semester1start) && !EmptyOrIsDate(semester1end) && 
      !EmptyOrIsDate(semester2start) && !EmptyOrIsDate(semester2end) && 
      !EmptyOrIsDate(semester3start) && !EmptyOrIsDate(semester3end)) {
       this.Ctx.ResponseWriter.WriteHeader(400)
       return
  }
  p := models.Plan {
    Id: bson.NewObjectId(),
    Name: name,
    Cost: cost,
    Market: market,
    Semester1Start: semester1start,
    Semester1End: semester1end,
    Semester2Start: semester2start,
    Semester2End: semester2end,
    Semester3Start: semester3start,
    Semester3End: semester3end,
    MarketingText1: marketingtext1,
    MarketingText2: marketingtext2,
    MarketingText3: marketingtext3,
  }
  pdao.Insert(p)
  planJson, err := json.Marshal(p)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  this.Ctx.WriteString(string(planJson))
}

func (this *APIController) GetUserPlan() {
  user_id := this.Ctx.Input.Param(":user_id")
  user, err := udao.FindById(user_id)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  plan, err := pdao.FindById(user.MealPlanId.Hex())
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  planJson, err := json.Marshal(plan)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  this.Ctx.WriteString(string(planJson))
}

func (this *APIController) EditPlan() {
  plan_id := this.Ctx.Input.Query(":plan_id")
  plan, err := pdao.FindById(plan_id)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  plan.Name = this.Ctx.Input.Query("name")
  plan.Cost, err = strconv.Atoi(this.Ctx.Input.Query("cost"))
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(400)
    return
  }
  plan.Market = this.Ctx.Input.Query("market")
  plan.Semester1Start = this.Ctx.Input.Query("semester1start")
  plan.Semester1End = this.Ctx.Input.Query("semester1end")
  plan.Semester2Start = this.Ctx.Input.Query("semester2start")
  plan.Semester2End = this.Ctx.Input.Query("semester2end")
  plan.Semester3Start = this.Ctx.Input.Query("semester3start")
  plan.Semester3End = this.Ctx.Input.Query("semester3end")
  plan.MarketingText1 = this.Ctx.Input.Query("marketingtext1")
  plan.MarketingText2 = this.Ctx.Input.Query("marketingtext2")
  plan.MarketingText3 = this.Ctx.Input.Query("marketingtext3")
  if (!EmptyOrIsDate(plan.Semester1Start) && !EmptyOrIsDate(plan.Semester1End) && 
      !EmptyOrIsDate(plan.Semester2Start) && !EmptyOrIsDate(plan.Semester2End) && 
      !EmptyOrIsDate(plan.Semester3Start) && !EmptyOrIsDate(plan.Semester3End)) {
       this.Ctx.ResponseWriter.WriteHeader(400)
       return
  }
  pdao.Update(plan)
  planJson, err := json.Marshal(plan)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  this.Ctx.WriteString(string(planJson))
}

func (this *APIController) DuplicatePlan() {
  plan_id := this.Ctx.Input.Query(":plan_id")
  plan, err := pdao.FindById(plan_id)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(400)
    return
  }
  new_plan := models.Plan {
    Id: bson.NewObjectId(),
    Name: plan.Name,
    Cost: plan.Cost,
    Market: plan.Market,
    Semester1Start: plan.Semester1Start,
    Semester1End: plan.Semester1End,
    Semester2Start: plan.Semester2Start,
    Semester2End: plan.Semester2End,
    Semester3Start: plan.Semester3Start,
    Semester3End: plan.Semester3End,
    MarketingText1: plan.MarketingText1,
    MarketingText2: plan.MarketingText2,
    MarketingText3: plan.MarketingText3,
  }
  pdao.Insert(new_plan)
  planJson, err := json.Marshal(new_plan)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  this.Ctx.WriteString(string(planJson))
}
