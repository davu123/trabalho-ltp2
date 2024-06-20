
2. **Arquivo models/review.go**

```go
package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
    ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    RestaurantID primitive.ObjectID `json:"restaurant_id" bson:"restaurant_id"`
    UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
    Rating       float32            `json:"rating" bson:"rating"`
    Comment      string             `json:"comment" bson:"comment"`
    CreatedAt    primitive.DateTime `json:"created_at" bson:"created_at"`
}
