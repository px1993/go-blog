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

func (this *AuthController) Register() {
        if "POST" == this.Ctx.Input.Method() {
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

        this.TplName = "auth/register.tpl" 
}

func (this *AuthController) Login() {
        //登录提交
        if "POST" == this.Ctx.Input.Method() {
                username := this.GetString("username")
                password := this.GetString("password")

                //查找用户
                user := models.User{Username:username}
                models.GetUsers(&user,"Username")
                if user.Password == password {
                        this.SetSession("userLogin",int64(user.Id))
                        this.Redirect("/user",302)
                        return
                }
                this.Redirect("/auth",302)
                return
        }

        this.TplName = "auth/login.tpl"
}

func (this *AuthController) Logout(){
        this.DelSession("userLogin")
        this.Redirect("/auth/login",302)
}