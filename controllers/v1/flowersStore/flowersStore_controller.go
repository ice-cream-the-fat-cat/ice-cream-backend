package flowersStore_controllers

import (
	"time"

	users_controllers "github.com/ice-cream-backend/controllers/v1/users"
	mongo_connection "github.com/ice-cream-backend/database"
	flowersStore_models "github.com/ice-cream-backend/models/v1/flowersStore"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func BuyNewFlower(flowersStore flowersStore_models.FlowersStore) (*mongo.UpdateResult, error) {
	ctx, ctxCancel := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)
	defer ctxCancel()

	collection := mongo_connection.MongoCollection(client, "users")

	userData, _ := users_controllers.GetUserByFireBaseUserId(flowersStore.FireBaseUserId)

	oid, _ := primitive.ObjectIDFromHex(flowersStore.FlowerID)

	updatedUser := bson.M{
		"$set": bson.M{
			"numCoins":          userData.NumCoins - flowersStore.Price,
			"flowerCollections": append(userData.FlowerCollections, oid),
			"lastUpdate":        time.Now(),
		},
	}
	result, updateErr := collection.UpdateByID(ctx, userData.ID, updatedUser)

	return result, updateErr
}
