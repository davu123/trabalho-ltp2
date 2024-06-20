package controllers

import (
	"context"
	"net/http"
	"restaurant-review/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func SearchRestaurants(c *gin.Context) {
	name := c.Query("name")
	address := c.Query("address")

	filter := bson.M{}
	if name != "" {
		filter["name"] = bson.M{"$regex": name, "$options": "i"}
	}
	if address != "" {
		filter["address"] = bson.M{"$regex": address, "$options": "i"}
	}

	var restaurants []models.Restaurant
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := restaurantCollection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var restaurant models.Restaurant
		cursor.Decode(&restaurant)
		restaurants = append(restaurants, restaurant)
	}

	c.JSON(http.StatusOK, restaurants)
}

func FilterRestaurants(c *gin.Context) {
	minRating, err := strconv.ParseFloat(c.Query("min_rating"), 32)
	if err != nil {
		minRating = 0
	}

	filter := bson.M{
		"rating": bson.M{"$gte": minRating},
	}

	var restaurants []models.Restaurant
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := restaurantCollection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var restaurant models.Restaurant
		cursor.Decode(&restaurant)
		restaurants = append(restaurants, restaurant)
	}

	c.JSON(http.StatusOK, restaurants)
}

