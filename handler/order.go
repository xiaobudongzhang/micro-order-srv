package handler

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/xiaobudongzhang/micro-order-srv/model/orders"
	proto "github.com/xiaobudongzhang/micro-order-srv/proto/order"
	"github.com/xiaobudongzhang/seata-golang/client"
	"github.com/xiaobudongzhang/seata-golang/client/at/exec"
	"github.com/xiaobudongzhang/seata-golang/client/at/sql/struct/cache"
	"github.com/xiaobudongzhang/seata-golang/client/config"
	context2 "github.com/xiaobudongzhang/seata-golang/client/context"
)

var (
	ordersService orders.Service
)

type Orders struct {
}

func Init() {
	ordersService, _ = orders.GetService()
}

func (e *Orders) New(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {

	rex,_:= metadata.FromContext(ctx)



	config.InitConf("D:\\micro\\micro-order-srv\\conf\\seate_client.yml")
	client.NewRpcClient()
	cache.SetTableMetaCache(cache.NewMysqlTableMetaCache(config.GetClientConfig().ATConfig.DSN))
	exec.InitDataResourceManager()

	rootContext := &context2.RootContext{Context:ctx}
	rootContext.Bind(rex["Xid"])

	fmt.Printf("rootContext:%v", rootContext)

	orderId, err := ordersService.New(req.BookId, req.UserId, req.OrderId, rootContext)

	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}

		return
	}

	rsp.Order = &proto.Order{
		Id: orderId,
	}

	return
}

//get order
func (e *Orders) GetOrder(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	log.Logf("[getorder] 收到获取订单请求 %d", req.OrderId)

	rsp.Order, err = ordersService.GetOrder(req.OrderId)

	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}
		return
	}

	rsp.Success = true
	return
}
