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
		router.POST("/registration", h.registrationAccount)
	}

	accounts := router.Group("/accounts")
	{
		accounts.GET("/:accountId", h.notAuthMiddleware, h.getAccount)
		accounts.GET("/search", h.notAuthMiddleware, h.getAccounts)
		accounts.PUT("/:accountId", h.requiredAuthMiddleware, h.updateAccount)
		accounts.DELETE("/:accountId", h.requiredAuthMiddleware, h.deleteAccount)
	}

	location := router.Group("/locations")
	{
		location.GET("/:pointId", h.notAuthMiddleware, h.getLocation)
		location.POST("", h.requiredAuthMiddleware, h.createLocation)
		location.PUT("/:pointId", h.requiredAuthMiddleware, h.updateLocation)
		location.DELETE("/:pointId", h.requiredAuthMiddleware, h.removeLocation)
	}

	animal := router.Group("/animals")
	{
		animal.GET("/:animalId", h.notAuthMiddleware, h.getAnimal)
		animal.GET("/search", h.notAuthMiddleware, h.getAnimals)
		animal.POST("", h.requiredAuthMiddleware, h.createAnimal)
		animal.PUT("/:animalId", h.requiredAuthMiddleware, h.updateAnimal)
		animal.DELETE("/:animalId", h.requiredAuthMiddleware, h.removeAnimal)

		animalID := animal.Group("/:animalId")
		{
			visitedLocation := animalID.Group("/locations")
			{
				visitedLocation.GET("", h.notAuthMiddleware, h.getVisitedLocation)
				visitedLocation.POST("/:pointId", h.requiredAuthMiddleware, h.createVisitedLocation)
				visitedLocation.PUT("", h.requiredAuthMiddleware, h.updateVisitedLocation)
				visitedLocation.DELETE("/:visitedPointId", h.requiredAuthMiddleware, h.removeVisitedLocation)
			}

			types := animalID.Group("/types")
			{
				types.POST("/:typeId", h.requiredAuthMiddleware, h.linkAnimalTypeToAnimal)
				types.PUT("", h.requiredAuthMiddleware, h.updateAnimalTypeToAnimal)
				types.DELETE("/:typeId", h.requiredAuthMiddleware, h.removeAnimalTypeToAnimal)
			}
		}

		animalType := animal.Group("/types")
		{
			animalType.GET("/:typeId", h.notAuthMiddleware, h.getAnimalType)
			animalType.POST("", h.requiredAuthMiddleware, h.createAnimalType)
			animalType.PUT("/:typeId", h.requiredAuthMiddleware, h.updateAnimalType)
			animalType.DELETE("/:typeId", h.requiredAuthMiddleware, h.removeAnimalType)
		}
	}

	return router
}
