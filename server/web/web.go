package web

import (
	"air/order/base"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/api/server/http"
	"github.com/gin-gonic/gin"
)


var router = gin.Default()


func Run() error {
	base.Micro.Init(
		micro.Name("order-web"),
		micro.Version("1.0.0"),
	)
	web := http.NewServer(":5066")
	auth.init()
	web.Handle("/", router)
	return web.Start()
}
