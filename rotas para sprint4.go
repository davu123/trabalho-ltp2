package routes

import (
    "github.com/gin-gonic/gin"
    "restaurant-review/controllers"
)

func RestaurantRoutes(router *gin.Engine) {
    router.POST("/restaurants", controllers.CreateRestaurant)
    router.GET("/restaurants/search", controllers.SearchRestaurants)
    router.GET("/restaurants/filter", controllers.FilterRestaurants)
    router.PUT("/restaurants/:id", controllers.UpdateRestaurant)
    router.DELETE("/restaurants/:id", controllers.DeleteRestaurant)
}
