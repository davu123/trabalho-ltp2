package main

import (
    "github.com/gin-gonic/gin"
    "restaurant-review/routes"
    "restaurant-review/database"
    "restaurant-review/config"
)

func main() {
    config.LoadConfig()
    database.ConnectDB()

    r := gin.Default()

    routes.AuthRoutes(r)
    routes.RestaurantRoutes(r)
    routes.ReviewRoutes(r)

    r.Run(":8080")
}
