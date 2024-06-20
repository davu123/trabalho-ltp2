package routes

import (
    "github.com/gin-gonic/gin"
    "restaurant-review/controllers"
)

func ReviewRoutes(router *gin.Engine) {
    router.POST("/reviews", controllers.CreateReview)
    router.GET("/reviews/:restaurant_id", controllers.GetReviews)
}
