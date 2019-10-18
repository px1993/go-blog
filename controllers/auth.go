package controllers

import (
        "github.com/astaxie/beego"
        "blog/models"
)

type AuthController struct {
        beego.Controller
}

func (this *AuthController) Prepare() {

}

var loginUrl = "/user"

func (this *AuthController) Get() {
        this.TplName = "auth/register.tpl" 
}


func (this *AuthController) Post(){
        //校验参数是否合法
        username := this.GetString("username")
        password   := this.GetString("password")
        repassword   := this.GetString("repassword")

        if password != repassword {
                this.Ctx.WriteString("两次密码输入不一致")
                return
        }

        //获取用户信息
        user := models.User{Username:username,Password:password}
        userId := models.AddUser(&user)

        this.SetSession("userLogin", userId)

        this.Redirect(loginUrl,302)
}
