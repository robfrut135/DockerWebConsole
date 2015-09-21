package main

import (
	_ "dockerwebconsole/models"
	_ "dockerwebconsole/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/auth"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", "acme.db")
	name := "default"
	force := false
	verbose := false
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	beego.SessionOn = true
	beego.InsertFilter("/appadmin/*", beego.BeforeRouter, auth.Basic("youradminname", "YourAdminPassword"))
	beego.Run()
}
