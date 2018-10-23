package routers

import (
	"mealplans/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/get_users/", &controllers.APIController{}, "get:GetAllUsers")
    beego.Router("/get_plans/", &controllers.APIController{}, "get:GetAllPlans")
    beego.Router("/get_plans/market/:market:string/", &controllers.APIController{}, "get:GetAllPlansInMarket")
    beego.Router("/get_plan/:plan_id:string/", &controllers.APIController{}, "get:GetPlan")
    beego.Router("/get_plan/user/:user_id:string/:token:string/", &controllers.APIController{}, "get:GetUserPlan")
    beego.Router("/edit_plan/:plan_id:string", &controllers.APIController{}, "post:EditPlan")
    beego.Router("/duplicate_plan/:plan_id:string", &controllers.APIController{}, "post:DuplicatePlan")
    //beego.Router("/create_user/", &controllers.APIController{}, "post:CreateUser")
    beego.Router("/populate_db/", &controllers.APIController{}, "get:PopulateDB")

    beego.Router("/create_plan/", &controllers.APIController{}, "post:CreatePlan")

    beego.InsertFilter("/create_plan/", beego.BeforeRouter, controllers.LoginFilter)
    beego.InsertFilter("/edit_plan/:plan_id:string", beego.BeforeRouter, controllers.LoginFilter)
    beego.InsertFilter("/get_plan/user/:user_id:string/:token:string/", beego.BeforeRouter, controllers.TokenFilter)
}
