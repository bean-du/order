package client

import (
	"air/order/base"
	auth "air/order/proto"
)

var AuthClient  auth.AuthService

func init()  {
	AuthClient = auth.NewAuthService("auth", base.Micro.Client())
}
