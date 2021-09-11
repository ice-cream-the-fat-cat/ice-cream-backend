package rules_models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rules struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	GardenId primitive.ObjectID `bson:"gardenId" json:"gardenId"`
	IsRemoved bool `bson:"isRemoved" json:"isRemoved"`
	CreatedDate time.Time `bson:"createdDate" json:"createdDate"`
	LastUpdate time.Time `bson:"lastUpdate" json:"lastUpdate"`
}

var RulesSchema = bson.M{
	"bsonType": "object",
	"required": []string{"name", "gardenId"},
	"additionalProperties": true,
	"properties": bson.M{
		"name": bson.M{
			"bsonType": "string",
			"description": "must be a string and is required",
		},
	},
}