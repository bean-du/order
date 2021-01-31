package service

import (
	"air/order/client"
	auth2 "air/order/proto"
	"golang.org/x/net/context"
)

type auth struct {}

var Auth = &auth{}

func (a *auth) Login(ctx context.Context, request *auth2.LoginRequest) (token string, err error) {
	login, err := client.AuthClient.Login(ctx, request)
	if err != nil {
		return "", err
	}
	return login.Token, nil
}
