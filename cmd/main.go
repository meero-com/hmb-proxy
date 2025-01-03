package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/meero-com/guild-proxy/pkg/api"
)

func main() {
	router := gin.New()
	fmt.Println("Hello World!")
	api.Activate(router)
	router.Run()
}
