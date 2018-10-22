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
    beego.Router("/get_plan/:id:string/", &controllers.APIController{}, "get:GetPlan")
    //beego.Router("/create_user/", &controllers.APIController{}, "post:CreateUser")
    beego.Router("/populate_db/", &controllers.APIController{}, "get:PopulateDB")
}
