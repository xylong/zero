// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: order.proto

package orderclient

import (
	"context"

	"zero/apps/order/rpc/order"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	OrderItem      = order.OrderItem
	OrdersRequest  = order.OrdersRequest
	OrdersResponse = order.OrdersResponse

	Order interface {
		Orders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) Orders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.Orders(ctx, in, opts...)
}
