package subscriber

import (
	"context"

	"github.com/xiaobudongzhang/micro-order-srv/model/orders"

	"github.com/micro/go-micro/v2/util/log"

	payS "github.com/xiaobudongzhang/micro-payment-srv/proto/payment"
)

var (
	orderService orders.Service
)

func Init() {
	orderService, _ = orders.GetService()
}

func PayOrder(ctx context.Context, event *payS.PayEvent) (err error) {

	log.Logf("收到支付订单通知, %d, %d", event.OrderId, event.State)

	err = orderService.UpdateOrderState(event.OrderId, int(event.State))

	if err != nil {
		log.Logf("收到支付单通知,更新状态异常, %s", err)
		return
	}
	return
}
