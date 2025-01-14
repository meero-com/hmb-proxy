package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/meero-com/guild-proxy/config"
	"github.com/meero-com/guild-proxy/pkg/api"
)

func main() {
	config.InitConfig()
	router := gin.New()
	fmt.Println("Hello World!")
	api.Activate(router)
	router.Run()
}
