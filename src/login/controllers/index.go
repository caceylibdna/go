package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
        "login/models"
)

type IndexController struct {
	beego.Controller
}

func (index *IndexController) Get() {
	sess := index.StartSession()
	username := sess.Get("username")
	fmt.Println(username)
	if username == nil || username == "" {
		index.TplName = "index.tpl"
	} else {
		index.TplName = "success.tpl"
	}

}

func (index *IndexController) Post() {
	sess := index.StartSession()
	var user models.User
	inputs := index.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("username")
	user.Pwd = inputs.Get("pwd")
        user.LoadId = inputs.Get("loadId")
        fmt.Println(user.Username)
        fmt.Println(user.LoadId)
        if user.LoadId == "Login" {
            err := models.ValidateUser(user)
            if err == nil {
	        sess.Set("username", user.Username)
                fmt.Println("username:", sess.Get("username"))
		index.TplName = "success.tpl"
            } else {
		index.TplName = "error.tpl"
	    }
        } else if user.LoadId == "Discover" {
            user.CreTyp = inputs.Get("creType")
            fmt.Println("dddddiscover")
            user.CrePassword = inputs.Get("crePassword")
            fmt.Println(user.CrePassword)
            user.CreUsername = inputs.Get("creUsername")
            user.NetName = inputs.Get("netName")
            user.Network = inputs.Get("network") 
            fmt.Println("discoverusername:", sess.Get("username"))
            err := models.AddCreNetworkToTb(user)
            fmt.Println("no error")
            if err == nil {
                index.TplName = "tasks.tpl"
            } else {
                index.TplName = "error.tpl"
            }
        } else {
            index.TplName = "error.tpl"
        }
}
