package controllers

import (
    "context"
    "net/http"
    "restaurant-review/models"
    "restaurant-review/database"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var reviewCollection *mongo.Collection = database.DB.Database("restaurantreview").Collection("reviews")

func CreateReview(c *gin.Context) {
    var input models.Review
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    input.ID = primitive.NewObjectID()
    input.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    _, err := reviewCollection.InsertOne(ctx, input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Review created successfully"})
}

func GetReviews(c *gin.Context) {
    restaurantID := c.Param("restaurant_id")
    objID, err := primitive.ObjectIDFromHex(restaurantID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid restaurant ID"})
        return
    }

    var reviews []models.Review
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := reviewCollection.Find(ctx, bson.M{"restaurant_id": objID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var review models.Review
        cursor.Decode(&review)
        reviews = append(reviews, review)
    }

    c.JSON(http.StatusOK, reviews)
}
