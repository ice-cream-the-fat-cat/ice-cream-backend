package users_controllers

import (
	"log"

	mongo_connection "github.com/ice-cream-backend/database"

	users_models "github.com/ice-cream-backend/models/v1/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(createdUserPost users_models.Users) (*mongo.InsertOneResult, error) {
	ctx := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)

	collection := mongo_connection.MongoCollection(client, "users")

	res, insertErr := collection.InsertOne(ctx, createdUserPost)

	if insertErr != nil {
		log.Println("Error creating new createGardens:", insertErr)
	}

	return res, insertErr
}

func GetUserByUserId(createUserId interface{}) (users_models.Users, error) {
	ctx := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)
	collection := mongo_connection.MongoCollection(client, "users")

	var result users_models.Users
	err := collection.FindOne(ctx, bson.D{
		primitive.E{Key: "_id", Value: createUserId},
	},
	).Decode(&result)

	if err != nil {
		log.Println("err in findOne:", err)
	}

	return result, err
}
