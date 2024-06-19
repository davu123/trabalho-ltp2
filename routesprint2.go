package routes

import (
	"restaurant-review/controllers"

	"github.com/gin-gonic/gin"
)

func RestaurantRoutes(router *gin.Engine) {
	router.POST("/restaurants", controllers.CreateRestaurant)
}

