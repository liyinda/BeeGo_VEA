package routers

import (
	"lfpgame/controllers"
	"github.com/astaxie/beego"
)

func init() {
    // 默认登录
    loginController := &controllers.LoginController{}
	beego.Router("/", loginController, "*:LoginIn")
	beego.Router("/login", loginController, "*:LoginIn")
	beego.Router("/userinfo", loginController, "*:UserInfo")
	beego.Router("/logout", loginController, "*:LoginOut")
	beego.Router("/no_auth", loginController, "*:NoAuth")

}
