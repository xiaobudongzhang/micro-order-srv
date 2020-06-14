package orders

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/xiaobudongzhang/micro-basic/common"
	"github.com/xiaobudongzhang/micro-plugins/db"
)

func (s *service) New(bookId int64, userId int64, hisId int64) (orderId int64, err error) {

	o := db.GetDB()
	insertSQL := `insert orders (user_id, book_id,inv_his_id,state) values (?,?,?,?)`

	r, err := o.Exec(insertSQL, userId, bookId, hisId, common.InventoryHistoryStateNotOut)

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
