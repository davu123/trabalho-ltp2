package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Address string             `json:"address" bson:"address"`
	Rating  float32            `json:"rating" bson:"rating"`
}

