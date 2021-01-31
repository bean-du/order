package main

import (
	"air/order/server/grpc"
	"air/order/server/web"
)

func main() {
	if err := web.Run(); err != nil {
		panic(err)
	}
	if err := grpc.Run(); err != nil {
		panic(err)
	}
}
