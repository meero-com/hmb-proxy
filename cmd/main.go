package main

import (
	"github.com/gin-gonic/gin"
	"github.com/meero-com/guild-proxy/pkg/api"
	"github.com/meero-com/guild-proxy/pkg/config"
)

func main() {
	config.InitConfig()
	config.PrintConfig()
	router := gin.New()
	api.Activate(router)
	router.Run()
}
