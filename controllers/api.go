package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func (u *ApiController) Post() {
	u.Data["json"] = map[string]string{"uid": "123"}
	u.ServeJSON()
}

func (u *ApiController) Get() {
	u.Data["json"] = "data json"
	u.ServeJSON()
}
