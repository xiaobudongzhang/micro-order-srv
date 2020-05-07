package main

import (
	"fmt"
	"micro-order-srv/handler"
	"micro-order-srv/subscriber"

	"github.com/micro-in-cn/tutorials/microservice-in-micro/part2/basic"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part2/basic/config"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part4/payment-srv/model"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part8/basic/common"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/api/resolver/micro"
	"github.com/micro/go-micro/v2/config/source/cli"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/xiaobudongzhang/micro-basic"

	order "micro-order-srv/proto/order"
)

func main() {
	basic.Init()

	micReg := etcd.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.service.order"),
		micro.Register(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(micro.Action(func(c *cli.Context) error {
		model.Init()
		handler.Init()
		subscriber.Init()

		return nil
	}))

	// Register Struct as Subscriber
	err := micro.RegisterSubscriber(common.TopicPaymentDone, service.Server(), subscriber.PayOrder)
	if err != nil {
		log.Fatal(err)
	}
	// Register Handler
	err = order.RegisterOrderHandler(service.Server(), new(handler.Order))
	if err != nil {
		log.Fatal(err)
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
