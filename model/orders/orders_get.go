package orders

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/xiaobudongzhang/micro-basic/db"
	proto "github.com/xiaobudongzhang/micro-order-srv/proto/order"
)

func (s *service) GetOrder(orderId int64) (order *proto.Order, err error) {
	order = &proto.Order{}

	o := db.GetDB()

	err = o.QueryRow("select id, user_id,book_id,inv_his_id,state from orders where id = ?", orderId).Scan(
		&order.Id, &order.UserId, &order.BookId, &order.InvHistoryId, &order.State
	)

	if err != nil {
		log.Logf("查询数据失败 err:%s", err)
		return
	}

	return
}
