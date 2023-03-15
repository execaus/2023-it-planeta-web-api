package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) sendCreatedWithBody(c *gin.Context, body interface{}) {
	c.JSON(http.StatusCreated, body)
}

func (h *Handler) sendOKWithBody(c *gin.Context, body interface{}) {
	c.JSON(http.StatusOK, body)
}

func (h *Handler) sendOk(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) sendBadRequest(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, message)
}

func (h *Handler) sendConflict(c *gin.Context) {
	c.AbortWithStatus(http.StatusConflict)
}

func (h *Handler) sendInternalServerError(c *gin.Context) {
	c.AbortWithStatus(http.StatusInternalServerError)
}

func (h *Handler) sendForbidden(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

func (h *Handler) sendNotFound(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotFound)
}

func (h *Handler) sendUnauthorized(c *gin.Context) {
	c.AbortWithStatus(http.StatusUnauthorized)
}
