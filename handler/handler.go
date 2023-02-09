package handler

import (
	"2023-it-planeta-web-api/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	{
		router.POST("/registration", notAuthMiddleware, h.registrationAccount)
	}

	accounts := router.Group("/accounts", requiredAuthMiddleware)
	{
		accounts.GET("/:accountId", h.getAccount)
		accounts.GET("/search", h.getAccounts)
	}

	return router
}

func requiredAuthMiddleware(c *gin.Context) {

}

func notAuthMiddleware(c *gin.Context) {

}
