package gardens_controllers

import (
	"log"

	mongo_connection "github.com/ice-cream-backend/database"
	gardens_models "github.com/ice-cream-backend/models/v1/gardens"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateGardens(createdGardensPost gardens_models.Gardens) (*mongo.InsertOneResult, error) {
	ctx := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)

	collection := mongo_connection.MongoCollection(client, "gardens")

	res, insertErr := collection.InsertOne(ctx, createdGardensPost)

	if insertErr != nil {
		log.Println("Error creating new createGardens:", insertErr)
	}

	return res, insertErr
}

func GetGardensById(createGardenId interface{}) gardens_models.Gardens {
	ctx := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)

	collection := mongo_connection.MongoCollection(client, "gardens")

	var result gardens_models.Gardens
	collection.FindOne(ctx, bson.D{
		primitive.E{Key: "_id", Value: createGardenId},
	},
	).Decode(&result)

	return result
}
