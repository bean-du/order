package service

import (
	"air/order/client"
	auth2 "air/order/proto"
	"air/order/schame"
	"golang.org/x/net/context"
)

type auth struct {
}

var Auth = &auth{}

func (a *auth) Login(ctx context.Context, request *schame.AuthRequest) (token string, err error) {
	login, err := client.AuthClient.Login(ctx, &auth2.LoginRequest{Username: request.Username, Password: request.Password})
	if err != nil {
		return "", err
	}
	return login.Token, nil
}
