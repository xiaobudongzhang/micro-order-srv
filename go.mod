module github.com/xiaobudongzhang/micro-order-srv

go 1.14

replace github.com/xiaobudongzhang/micro-basic => /wwwroot/microdemo/micro-basic

replace github.com/xiaobudongzhang/micro-inventory-srv => /wwwroot/microdemo/micro-inventory-srv

replace github.com/xiaobudongzhang/micro-payment-srv => /wwwroot/microdemo/micro-payment-srv

replace github.com/xiaobudongzhang/micro-order-srv => /wwwroot/microdemo/micro-order-srv

replace github.com/xiaobudongzhang/micro-plugins => /wwwroot/microdemo/micro-plugins

require (
	github.com/go-log/log v0.2.0
	github.com/golang/protobuf v1.4.1
	github.com/micro-in-cn/tutorials/microservice-in-micro v0.0.0-20200430044506-2451e30bf530
	github.com/micro/go-micro/v2 v2.6.0
	github.com/xiaobudongzhang/micro-basic v1.1.5
	github.com/xiaobudongzhang/micro-inventory-srv v1.0.0
	github.com/xiaobudongzhang/micro-payment-srv v0.0.0-00010101000000-000000000000
	github.com/xiaobudongzhang/micro-plugins v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.26.0
	google.golang.org/protobuf v1.22.0
)
