package controllers

import (
  "strconv"
	"github.com/astaxie/beego"
  "encoding/json"
  "mealplans/models"
  "gopkg.in/mgo.v2/bson"
)

type PlanAPIController struct {
  beego.Controller
}

func (this *PlanAPIController) GetPlan() {
  id := this.Ctx.Input.Param(":plan_id")
  if !bson.IsObjectIdHex(id) {
    this.Ctx.ResponseWriter.WriteHeader(400)
    return
  }
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

func (this *PlanAPIController) GetAllPlans() {
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

func (this *PlanAPIController) GetAllPlansInMarket() {
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

func (this *PlanAPIController) CreatePlan() {
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
func (this *PlanAPIController) GetUserPlan() {
  user_id := this.Ctx.Input.Param(":user_id")
  if !bson.IsObjectIdHex(user_id) {
    this.Ctx.ResponseWriter.WriteHeader(400)
    return
  }
  user, err := udao.FindById(user_id)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(500)
    return
  }
  plan, err := pdao.FindById(user.MealPlanId)
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

func (this *PlanAPIController) EditPlan() {
  plan_id := this.Ctx.Input.Query("plan_id")
  if !bson.IsObjectIdHex(plan_id) {
    this.Ctx.ResponseWriter.WriteHeader(400)
    return
  }
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

func (this *PlanAPIController) DuplicatePlan() {
  plan_id := this.Ctx.Input.Query("plan_id")
  if !bson.IsObjectIdHex(plan_id) {
    this.Ctx.ResponseWriter.WriteHeader(400)
    return
  }
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

func (this *PlanAPIController) DeletePlan() {
  plan_id := this.Ctx.Input.Query("plan_id")
  if !bson.IsObjectIdHex(plan_id) {
    this.Ctx.ResponseWriter.WriteHeader(400)
    return
  }
  plan, err := pdao.FindById(plan_id)
  if err != nil {
    this.Ctx.ResponseWriter.WriteHeader(400)
    return
  }
  users, err := udao.FindByMealPlanId(plan_id)
  for _, user := range users {
    user.MealPlanId = ""
    udao.Update(user)
  }
  pdao.Delete(plan)
  this.Ctx.WriteString("Plan deleted.")
}
