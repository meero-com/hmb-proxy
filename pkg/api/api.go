package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// awsSdk "github.com/meero-com/guild-proxy/pkg/aws"
)

type handler struct{}
type payload struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func Activate(router *gin.Engine) {
	newHandler(router)
}

func newHandler(router *gin.Engine) {
	h := handler{}

	router.GET("/api/test", h.Get)
	router.POST("/api/test", h.Create)
}

func (h *handler) Get(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World from api!")
}

func (h *handler) Create(c *gin.Context) {
	ch := make(chan string)
	var content payload
	if err := c.ShouldBindJSON(&content); err != nil {
		log.Println(err)
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	go process(ch, content)
	r := <-ch
	c.JSON(http.StatusOK, gin.H{
		"anwser": r,
	})
}
