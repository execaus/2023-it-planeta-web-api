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
		location.POST("", h.createLocation)
		location.PUT("/:pointId", h.updateLocation)
		location.DELETE("/:pointId", h.removeLocation)
	}

	animal := router.Group("/animals")
	{
		visitedLocation := animal.Group("/:animalId/locations")
		{
			// visitedLocation.GET("", h.getVisitedLocation)
			visitedLocation.POST("/:pointId", h.createVisitedLocation)
			visitedLocation.PUT("", h.updateVisitedLocation)
			visitedLocation.DELETE("/:visitedPointId", h.removeVisitedLocation)
		}

		animal.GET("/:animalId", h.getAnimal)

		animalType := animal.Group("/types")
		{
			animalType.GET("/:typeId", h.getAnimalType)
			animalType.POST("", h.createAnimalType)
			animalType.PUT("/:typeId", h.updateAnimalType)
			animalType.DELETE("/:typeId", h.removeAnimalType)
		}
	}

	return router
}
