package main

import (
	"fmt"
	"github.com/xiaobudongzhang/micro-order-srv/handler"
	"github.com/xiaobudongzhang/micro-order-srv/subscriber"
	"net/http"

	"github.com/xiaobudongzhang/micro-basic"
	"github.com/xiaobudongzhang/micro-basic/config"
	"github.com/xiaobudongzhang/micro-payment-srv/model"
	"github.com/xiaobudongzhang/micro-basic/common"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/xiaobudongzhang/micro-basic"

	order "micro-order-srv/proto/order"

	"github.com/micro/go-micro/v2"
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
	if err := service.Init(
		web.Action(
			func (c *cli.Context)  {
				model.Init()
				handler.Init()
				subscriber.Init()
				return nil
			})
	); err != nil {
		log.Fatal(err)
	}

	err := micro.RegisterSubscriber(common.TopicPaymentDone, service.Server(), subscriber.PayOrder)
	if err != nil {
		log.Fatal(err)
	}

	err = proto.RegisterOrdersHandler(service.Server(), new(handler.Orders))
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
