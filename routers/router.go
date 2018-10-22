package routers

import (
	"mealplans/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/get_users/", &controllers.APIController{}, "get:GetAllUsers")
}
