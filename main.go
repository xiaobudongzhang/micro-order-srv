package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"micro-order-srv/handler"
	"micro-order-srv/subscriber"

	order "micro-order-srv/proto/order"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.service.order"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	order.RegisterOrderHandler(service.Server(), new(handler.Order))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("mu.micro.book.service.order", service.Server(), new(subscriber.Order))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
