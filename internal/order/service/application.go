package service

import (
	"context"
	"github.com/mfzzf/order_microservices/common/metrics"
	"github.com/mfzzf/order_microservices/order/adapters"
	"github.com/mfzzf/order_microservices/order/app"
	"github.com/mfzzf/order_microservices/order/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(context context.Context) app.Application {
	orderRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			GetCustomerOrder: query.NewGetCustomerOrderHandler(orderRepo, logger, metricClient),
		},
	}
}
