package service

import (
	"context"
	"github.com/mfzzf/order_microservices/stock/app"
)

func NewApplication(context context.Context) app.Application {
	return app.Application{}
}
