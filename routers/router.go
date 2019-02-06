package routers

import (
	"myCMS/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/Admin/UserInfo/Index",&controllers.UserInfoController{},"get:Index")
}
