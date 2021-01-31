package grpc

import (
	"air/order/base"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
)

func Run() error {
	base.Micro.Init(
		micro.Name("order"),
		micro.Version("1.0.0"),
	)
	if err := base.Micro.Server().Init(server.Wait(nil)); err != nil {
		return err
	}
	return base.Micro.Run()
}
