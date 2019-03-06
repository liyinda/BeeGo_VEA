package main

import (
	_ "lfpgame/routers"
	"github.com/astaxie/beego"
	"lfpgame/models"
)

func main() {
	models.Init()
	beego.Run()
}

