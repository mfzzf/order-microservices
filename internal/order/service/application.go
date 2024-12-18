package service

import (
	"context"
	"github.com/mfzzf/order_microservices/order/adapters"
	"github.com/mfzzf/order_microservices/order/app"
)

func NewApplication(context context.Context) app.Application {
	orderRepo := adapters.NewMemoryOrderRepository()
	return app.Application{}
}
