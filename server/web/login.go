package web

import (
	auth1 "air/order/proto"
	"air/order/service"
	"air/order/utils"

	"github.com/gin-gonic/gin"
)

type orderAuth struct {
}

var auth = &orderAuth{}

func (o *orderAuth) init() {
	router.POST("/login", o.login)
}

func (o *orderAuth) login(c *gin.Context) {
	request := new(auth1.LoginRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		utils.HttpResponse(c, 1, "params error", nil)
		return
	}

	login, err := service.Auth.Login(c.Request.Context(), request)
	if err != nil {
		utils.HttpResponse(c, 1, "internal error", nil)
		return
	}
	utils.HttpResponse(c, 0, "ok", map[string]string{"token": login})
}
