// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	homestayOrder "zero-look/app/order/cmd/api/internal/handler/homestayOrder"
	"zero-look/app/order/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 创建民宿订单
				Method:  http.MethodPost,
				Path:    "/homestayOrder/createHomestayOrder",
				Handler: homestayOrder.CreateHomestayOrderHandler(serverCtx),
			},
			{
				// 用户订单明细
				Method:  http.MethodPost,
				Path:    "/homestayOrder/userHomestayOrderDetail",
				Handler: homestayOrder.UserHomestayOrderDetailHandler(serverCtx),
			},
			{
				// 用户订单列表
				Method:  http.MethodPost,
				Path:    "/homestayOrder/userHomestayOrderList",
				Handler: homestayOrder.UserHomestayOrderListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/order/v1"),
	)
}
