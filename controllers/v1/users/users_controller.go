package users_controllers

import (
	"log"
	"time"

	mongo_connection "github.com/ice-cream-backend/database"

	users_models "github.com/ice-cream-backend/models/v1/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(createdUserPost users_models.Users) (*mongo.InsertOneResult, error) {
	ctx, ctxCancel := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)
	defer ctxCancel()

	collection := mongo_connection.MongoCollection(client, "users")

	res, insertErr := collection.InsertOne(ctx, createdUserPost)

	if insertErr != nil {
		log.Println("Error creating new createGardens:", insertErr)
	}

	return res, insertErr
}

func GetUserByFireBaseUserId(createUserId interface{}) (users_models.Users, error) {
	ctx, ctxCancel := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)
	defer ctxCancel()

	collection := mongo_connection.MongoCollection(client, "users")

	var result users_models.Users
	err := collection.FindOne(ctx, bson.D{
		primitive.E{Key: "fireBaseUserId", Value: createUserId},
	},
	).Decode(&result)

	if err != nil {
		log.Println("err in findOne:", err)
	}

	return result, err
}

func UpdateUserByUserId(userId interface{}, user users_models.Users) (*mongo.UpdateResult, error) {
	ctx, ctxCancel := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)
	defer ctxCancel()

	collection := mongo_connection.MongoCollection(client, "users")

	updatedUser := bson.M{
		"$set": bson.M{
			"numCoins": user.NumCoins,
			"flowerCollections": user.FlowerCollections,
			"lastUpdate": time.Now(),
		},
	}

	result, updateErr := collection.UpdateByID(ctx, userId, updatedUser)

	return result, updateErr
}