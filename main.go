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

    r.Run(":8080")
}
