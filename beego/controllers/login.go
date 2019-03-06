/**********************************************
** @Des: login
** @Author: copy
** @Date:   2019-2-20
** @Last Modified by:   bjlfp
** @Last Modified time: 2019-2-20
***********************************************/
package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"encoding/json"

	"github.com/astaxie/beego"
	"lfpgame/libs"
	"lfpgame/models"
)

type LoginController struct {
	BaseController
}
type User struct {
	LoginName  string
	Password   string
}
//登录 TODO:XSRF过滤
func (self *LoginController) LoginIn() {
	if self.userId > 0 {
		self.redirect(beego.URLFor("HomeController.Index"))
	}
	errorMsg := ""
	authkey := ""
	if self.isPost() {
		var user User
		//fmt.Println(self.Ctx.Input.RequestBody)
		data := self.Ctx.Input.RequestBody
		//json data into user object
		err := json.Unmarshal(data, &user)
		if err != nil {
			fmt.Println("json.Unmarshal is err:", err.Error())
		}
		
		//fmt.Println(user)

		username := strings.TrimSpace(user.LoginName)
		password := strings.TrimSpace(user.Password)
		//fmt.Println(password)

		if username != "" && password != "" {
			user, err := models.AdminGetByName(username)
			fmt.Println(user)
			//flash := beego.NewFlash()
			
			if err != nil || user.Password != libs.Md5([]byte(password+user.Salt)) {
				errorMsg = "帐号或密码错误"
			} else if user.Status == -1 {
				errorMsg = "该帐号已禁用"
			} else {
				user.LastIp = self.getClientIp()
				user.LastLogin = time.Now().Unix()
				user.Update()
				authkey = libs.Md5([]byte(self.getClientIp() + "|" + user.Password + user.Salt))
				self.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)
				errorMsg = "ok"
				
				//fmt.Println(self.Data["json"])
				//self.redirect(beego.URLFor("HomeController.Index"))
			}
			
			//flash.Error(errorMsg)
			//flash.Store(&self.Controller)
			//self.redirect(beego.URLFor("LoginController.LoginIn"))
		}
	}
	self.Data["json"] = map[string]interface{}{"status": 200, "message": errorMsg, "token": authkey}
	fmt.Println(self.Data["json"])
    self.ajaxMsg(self.Data["json"] , 1)
	//self.TplName = "login/login.html"
}

//登出
func (self *LoginController) LoginOut() {
	self.Ctx.SetCookie("auth", "")
	self.Data["json"] = map[string]interface{}{"status": 200, "message": "loingOut", "token": ""}
	fmt.Println(self.Data["json"])
    self.ajaxMsg(self.Data["json"] , 1)
	//self.redirect(beego.URLFor("LoginController.LoginIn"))
}

func (self *LoginController) NoAuth() {
	self.Ctx.WriteString("没有权限")
}

// userinfo get
func (self *LoginController) UserInfo() {
	if self.userId > 0 {
		self.redirect(beego.URLFor("HomeController.Index"))
	}
	
	self.Data["json"] = map[string]interface{}{"roles": "['admin']", "name":"Super Admin" , "introduction": "我是超级管理员", "token": "admin" , "avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",}
	fmt.Println(self.Data["json"])
    self.ajaxMsg(self.Data["json"] , 1)
	//self.TplName = "login/login.html"
}