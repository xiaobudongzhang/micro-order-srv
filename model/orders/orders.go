package orders

import (
	"fmt"
	"github.com/xiaobudongzhang/seata-golang/client/context"
	"sync"

	"github.com/micro/go-micro/v2/client"
	invS "github.com/xiaobudongzhang/micro-inventory-srv/proto/inventory"
	proto "github.com/xiaobudongzhang/micro-order-srv/proto/order"
)

var (
	s         *service
	invClient invS.InventoryService
	m         sync.RWMutex
)

type service struct {
}

type Service interface {
	New(bookId int64, userId int64, hisId int64, ctx *context.RootContext) (orderId int64, err error)

	GetOrder(orderId int64) (order *proto.Order, err error)

	UpdateOrderState(orderId int64, state int) (err error)
}

func GetService() (Service, error) {

	if s == nil {
		return nil, fmt.Errorf("get service 未初始化")
	}
	return s, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}
	invClient = invS.NewInventoryService("mu.micro.book.service.inventory", client.DefaultClient)
	s = &service{}
}
