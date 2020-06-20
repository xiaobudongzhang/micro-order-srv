package orders

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/xiaobudongzhang/micro-basic/common"
	"github.com/xiaobudongzhang/micro-plugins/db"
	"github.com/xiaobudongzhang/seata-golang/client"

	"github.com/xiaobudongzhang/seata-golang/client/at/sql/struct/cache"
	"github.com/xiaobudongzhang/seata-golang/client/config"
	"github.com/xiaobudongzhang/seata-golang/client/context"

	"github.com/xiaobudongzhang/seata-golang/client/at/exec"
)

func (s *service) New(bookId int64, userId int64, hisId int64, ctx2 *context.RootContext) (orderId int64, err error) {



	config.InitConf("D:\\micro\\micro-order-srv\\conf\\seate_client.yml")
	client.NewRpcClient()
	cache.SetTableMetaCache(cache.NewMysqlTableMetaCache(config.GetClientConfig().ATConfig.DSN))
	exec.InitDataResourceManager()

	db,err := exec.NewDB(config.GetClientConfig().ATConfig)
	if err != nil {
		panic(err)
	}
	tx2, _ := db.Begin(ctx2)
	//o := db.GetDB()
	insertSQL := `insert micro_book_mall.orders (user_id, book_id,inv_his_id,state) values (?,?,?,?)`

	r, err := tx2.Exec(insertSQL, userId, bookId, hisId, common.InventoryHistoryStateNotOut)

	if err != nil {
		log.Logf("新增订单失败, err:%s", err)
		tx2.Rollback()

		return
	}
	tx2.Commit()

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
