package orders

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/xiaobudongzhang/micro-basic/common"
	"github.com/xiaobudongzhang/micro-basic/db"
	invS "github.com/xiaobudongzhang/micro-inventory-srv/proto/inventory"
)

func (s *service) New(bookId int64, userId int64) (orderId int64, err error) {
	rsp, err := invClient.Sell(content.TODO(), &invS.Request{
		BookId: bookId, UserId: userId,
	})

	if err != nil {
		log.Logf("sell 调用库存服务失败：%s", err.Error())
		return
	}

	o := db.GetDB()
	insertSQL := `insert orders (user_id, book_id,inv_his_id,state) values (?,?,?,?)`

	r, err := o.Exec(insertSQL, userId, bookId, rsp.InvH.Id, common.InventoryHistoryStateNotOut)

	if err != nil {
		log.Logf("新增订单失败, err:%s", err)
		return
	}

	orderId, _ = r.LastInsertId()
	return
}

func (s *service) UpdateOrderState(orderId int64, state int) (err error) {
	updateSQL := `update orders set state = ? where id = ? `

	o := db.GetDB()

	_, err = o.Exec(updateSQL, state, orderId)

	if err != nil {
		log.Logf("更新失败， err:%s", err)
		return
	}
	return
}
