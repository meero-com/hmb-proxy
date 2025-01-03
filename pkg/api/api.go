package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct{}

func Activate(router *gin.Engine) {
	newHandler(router)
}

func newHandler(router *gin.Engine) {
	h := handler{}

	router.GET("/api/test", h.Get)
}

func (h *handler) Get(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World from api!")
}
