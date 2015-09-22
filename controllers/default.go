package controllers

import (
	"dockerwebconsole/config"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type MainController struct {
	beego.Controller
	ConfigData config.Config
}

var FilterLogin = func(ctx *context.Context) {

	//******** This page requires login
	sess := ctx.Input.Session("acme")
	if sess == nil {
		ctx.Redirect(302, "/user/login/main")
		return
	}
	m := sess.(map[string]interface{})
	fmt.Println("username is", m["username"])
	fmt.Println("logged in at", m["timestamp"])
}

func (this *MainController) activeContent(view string) {
	this.Layout = "basic-layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.tpl"
	this.TplNames = view + ".tpl"
	this.Data["domainname"] = "localhost:8080"
	//this.Data["domainname"] = "yourdomainname"

	sess := this.GetSession("acme")
	if sess != nil {
		this.Data["InSession"] = 1 // for login bar in header.tpl
		m := sess.(map[string]interface{})
		this.Data["First"] = m["first"]
	}
}

func (this *MainController) Get() {
	this.activeContent("index")
}

func (this *MainController) About() {
	this.activeContent("about")
}

func (this *MainController) Main() {
	this.activeContent("main")

	//******** This page requires login
	sess := this.GetSession("acme")
	if sess == nil {
		this.Redirect("/user/login/main", 302)
		return
	}
	m := sess.(map[string]interface{})
	fmt.Println("username is", m["username"])
	fmt.Println("logged in at", m["timestamp"])
}

func (this *MainController) Notice() {
	this.activeContent("notice")

	flash := beego.ReadFromRequest(&this.Controller)
	if n, ok := flash.Data["notice"]; ok {
		this.Data["notice"] = n
	}
}
