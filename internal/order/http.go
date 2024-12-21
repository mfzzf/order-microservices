package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mfzzf/order_microservices/order/app"
	"github.com/mfzzf/order_microservices/order/app/query"
	"net/http"
)

type HTTPServer struct {
	app app.Application
}

func (H HTTPServer) PostCustomerCustomerIDOrders(c *gin.Context, customerID string) {
	//TODO implement me
	panic("implement me")
}

func (H HTTPServer) GetCustomerCustomerIDOrdersOrderID(c *gin.Context, customerID string, orderID string) {
	o, err := H.app.Queries.GetCustomerOrder.Handle(c, query.GetCustomerOrder{
		CustomerID: "fake-customer-id",
		OrderID:    "fake-ID",
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": o})
	}
}
