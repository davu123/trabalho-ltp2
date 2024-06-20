package routes

import (
	"restaurant-review/controllers"

	"github.com/gin-gonic/gin"
)

func RestaurantRoutes(router *gin.Engine) {
	router.POST("/restaurants", controllers.CreateRestaurant)
	router.GET("/restaurants/search", controllers.SearchRestaurants)
	router.GET("/restaurants/filter", controllers.FilterRestaurants)
	router.PUT("/restaurants/:id", controllers.UpdateRestaurant)
	router.DELETE("/restaurants/:id", controllers.DeleteRestaurant)
}

