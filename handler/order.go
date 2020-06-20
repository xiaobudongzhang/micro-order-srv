package handler

import (
	"context"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/xiaobudongzhang/micro-order-srv/model/orders"
	proto "github.com/xiaobudongzhang/micro-order-srv/proto/order"
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

	rootContext := &context2.RootContext{Context:ctx}
	rootContext.Bind(rex["Xid"])



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
