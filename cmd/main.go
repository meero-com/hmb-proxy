package main

import (
	"github.com/gin-gonic/gin"
	"github.com/meero-com/hmb-proxy/pkg/api"
	"github.com/meero-com/hmb-proxy/pkg/config"
)

func main() {
	config.InitConfig()
	config.PrintConfig()
	router := gin.New()
	api.Activate(router)
	router.Run()
}
