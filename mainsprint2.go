package main

import (
	"restaurant-review/config"
	"restaurant-review/database"
	"restaurant-review/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.ConnectDB()

	r := gin.Default()

	routes.AuthRoutes(r)
	routes.RestaurantRoutes(r)

	r.Run(":8080")
}

