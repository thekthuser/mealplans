package routers

import (
	"mealplans/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/get_users/", &controllers.UserAPIController{}, "get:GetAllUsers")
    beego.Router("/create_user/", &controllers.UserAPIController{}, "post:CreateUser")

    beego.Router("/get_plans/", &controllers.PlanAPIController{}, "get:GetAllPlans")
    beego.Router("/get_plans/market/:market:string/", &controllers.PlanAPIController{}, "get:GetAllPlansInMarket")
    beego.Router("/get_plan/:plan_id:string/", &controllers.PlanAPIController{}, "get:GetPlan")
    beego.Router("/get_plan/user/:user_id:string/:token:string/", &controllers.PlanAPIController{}, "get:GetUserPlan")

    beego.Router("/create_plan/", &controllers.PlanAPIController{}, "post:CreatePlan")
    beego.Router("/edit_plan/", &controllers.PlanAPIController{}, "post:EditPlan")
    beego.Router("/duplicate_plan/", &controllers.PlanAPIController{}, "post:DuplicatePlan")
    beego.Router("/delete_plan/", &controllers.PlanAPIController{}, "post:DeletePlan")

    beego.Router("/populate_db/", &controllers.APIController{}, "get:PopulateDB")


    beego.InsertFilter("/create_user/", beego.BeforeRouter, controllers.LoginFilter)
    beego.InsertFilter("/create_plan/", beego.BeforeRouter, controllers.LoginFilter)
    beego.InsertFilter("/delete_plan/", beego.BeforeRouter, controllers.LoginFilter)
    beego.InsertFilter("/edit_plan/:plan_id:string", beego.BeforeRouter, controllers.LoginFilter)
    beego.InsertFilter("/duplicate_plan/:plan_id:string", beego.BeforeRouter, controllers.LoginFilter)
    beego.InsertFilter("/get_plan/user/:user_id:string/:token:string/", beego.BeforeRouter, controllers.TokenFilter)
}
