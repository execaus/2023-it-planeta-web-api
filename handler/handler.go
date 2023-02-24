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
		accounts.PUT("/:accountId", h.updateAccount)
		accounts.DELETE("/:accountId", h.deleteAccount)
	}

	location := router.Group("/locations")
	{
		location.GET("/:pointId", h.getLocation)
	}

	animals := router.Group("/animals")
	{
		animals.GET("/:animalId", h.getAnimal)
	}

	return router
}
