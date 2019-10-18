package models

import (
        "github.com/astaxie/beego/orm"
        _"github.com/go-sql-driver/mysql"
)

type User struct{
        Id int
        Username string `orm:"size(100)"`
        Password string `orm:"size(255)"`
}

func init(){
        orm.RegisterDataBase("default","mysql","root:Px123456789@tcp(148.70.62.99:3306)/users?charset=utf8",30)
        orm.RegisterModel(new(User))
        orm.RunSyncdb("default", false, true)
}

//新增用户
func AddUser(user *User)(int64){
        o := orm.NewOrm()
        id,err := o.Insert(user)

        if err == nil {
                return id
        }
        return 0
}

//获取用户列表
func GetUsers(user *User,tableName string){
        o := orm.NewOrm()
        if tableName == "" {
                o.Read(user)
        } else {
                o.Read(user,tableName)
        }
}
