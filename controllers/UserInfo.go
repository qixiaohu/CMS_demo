package controllers

import "github.com/astaxie/beego"

type UserInfoController struct {
	beego.Controller
}

func (this *UserInfoController)Index(){
	this.TplName = "UserInfo/index.html"
}
