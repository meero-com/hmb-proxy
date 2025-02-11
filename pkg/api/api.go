package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type handler struct{}

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
	var content requestPayload

	if err := c.ShouldBindJSON(&content); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println("Failed to bind json error: %s", err)
		return
	}

	go process(ch, content)

	select {
	case r := <-ch:
		var out map[string]interface{}
		json.Unmarshal([]byte(r), &out)
		c.JSON(http.StatusOK, out)
	case <-time.After(time.Duration(content.Payload.Timeout) * time.Second):
		log.Println("requested timed out")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "backend service timed out",
		})
	}

}
