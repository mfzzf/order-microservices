package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RunHTTPServer(ServiceName string, wrapper func(router *gin.Engine)) {
	addr := viper.Sub(ServiceName).GetString("http-addr")
	if addr == "" {
		// TODO:Warning Log
	}
	RunHTTPServerOnAddr(addr, wrapper)
}

func RunHTTPServerOnAddr(addr string, wrapper func(router *gin.Engine)) {
	apiRouter := gin.New()
	wrapper(apiRouter)
	apiRouter.Group("/api")

	if err := apiRouter.Run(); err != nil {
		panic(err)
	}
}
