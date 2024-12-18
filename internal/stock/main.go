package main

import (
	"context"
	"github.com/mfzzf/order_microservices/common/config"
	"github.com/mfzzf/order_microservices/common/genproto/stockpb"
	"github.com/mfzzf/order_microservices/common/server"
	"github.com/mfzzf/order_microservices/stock/ports"
	"github.com/mfzzf/order_microservices/stock/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("stock.service-name")
	serviceType := viper.GetString("stock.service-to-run")
	log.Println(serviceType)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	application := service.NewApplication(ctx)
	switch serviceType {
	case "grpc":
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer(application)
			stockpb.RegisterStockServiceServer(server, svc)

		})
	case "http":
		//暂时不用
	default:
		panic("no serviceType set")
	}

}
