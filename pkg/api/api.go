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

	// TODO: consider variabilizing prefix
	router.GET("/api/health", h.Healthcheck)
	router.POST("/api/process", h.Create)
}

func (h *handler) Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
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
