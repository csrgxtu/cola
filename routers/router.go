package routers

import (
	"cola/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{}, "get:Welcome")
}
