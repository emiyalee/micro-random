package main

import (
	"math/rand"
	"time"

	"github.com/astaxie/beego"

	"github.com/emiyalee/micro-random/cmd/api-gateway/controllers"
	"github.com/emiyalee/micro-random/cmd/api-gateway/models"
	_ "github.com/emiyalee/micro-random/cmd/api-gateway/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	client, err := models.NewRandomClient()
	if err != nil {
		return
	}

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/random",
			beego.NSInclude(
				&controllers.RandomController{
					Client:    client,
					RandomNum: rand.Int31n(int32(time.Now().Nanosecond())),
				},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.Run()
}
