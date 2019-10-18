package controllers

import (
        "github.com/astaxie/beego"
        "blog/models"
)

var loginUserId int

type UserController struct{
        beego.Controller
}

func (this *UserController) Prepare(){
        userId := this.GetSession("userLogin")
        
        if userId == nil {
                this.Redirect("/auth/login",302)
        } else {
                loginUserId = int(userId.(int64))
        }

}

func (this *UserController) Get(){

        user := models.User{Id:loginUserId}
        models.GetUsers(&user,"")

        this.Data["userId"]= user.Id;
        this.Data["username"] = user.Username;
        this.Data["password"] = user.Password;

        this.TplName = "user/index.tpl"
}
