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
		animal.GET("/:animalId", h.getAnimal)
		animal.GET("/search", h.getAnimals)
		animal.POST("", h.createAnimal)
		animal.PUT("/:animalId", h.updateAnimal)
		animal.DELETE("/:animalId", h.removeAnimal)

		animalID := animal.Group("/:animalId")
		{
			visitedLocation := animalID.Group("/locations")
			{
				visitedLocation.GET("", h.getVisitedLocation)
				visitedLocation.POST("/:pointId", h.createVisitedLocation)
				visitedLocation.PUT("", h.updateVisitedLocation)
				visitedLocation.DELETE("/:visitedPointId", h.removeVisitedLocation)
			}

			types := animalID.Group("/types")
			{
				types.POST("/:typeId", h.linkAnimalTypeToAnimal)
				// types.PUT("", h.updateAnimalTypeInAnimal)
				// types.DELETE("/:typeId", h.removeAnimalTypeInAnimal)
			}
		}

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
