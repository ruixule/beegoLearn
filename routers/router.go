package routers

import (
	"github.com/2017RXL/beegoLearn/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/editor", &controllers.EditorController{})

}
