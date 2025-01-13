package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/meero-com/guild-proxy/pkg/api"
	"github.com/meero-com/guild-proxy/config"
)

func main() {
	config.InitConfig()
	router := gin.New()
	fmt.Println("Hello World!")
	api.Activate(router)
	router.Run()
}
