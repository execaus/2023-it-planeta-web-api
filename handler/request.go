package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) sendBadRequest(c *gin.Context) {
	c.AbortWithStatus(http.StatusBadRequest)
}

func (h *Handler) sendConflict(c *gin.Context) {
	c.AbortWithStatus(http.StatusConflict)
}

func (h *Handler) sendInternalServerError(c *gin.Context) {
	c.AbortWithStatus(http.StatusInternalServerError)
}

func (h *Handler) sendOK(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) sendOKWithBody(c *gin.Context, body interface{}) {
	c.JSON(http.StatusOK, body)
}
