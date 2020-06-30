package routers

import (
	"beegoLearn/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/editor", &controllers.EditorController{})

}
