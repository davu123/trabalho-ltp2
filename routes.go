package routes

import (
    "github.com/gin-gonic/gin"
    "restaurant-review/controllers"
)

func AuthRoutes(router *gin.Engine) {
    router.POST("/register", controllers.Register)
    router.POST("/login", controllers.Login)
}
