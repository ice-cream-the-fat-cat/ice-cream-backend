package garden_categories_controllers

import (
	"context"
	"log"

	mongo_connection "github.com/ice-cream-backend/database"
	garden_categories_models "github.com/ice-cream-backend/models/v1/garden_categories"
	"go.mongodb.org/mongo-driver/bson"
)

func GetGardenCategories() ([]garden_categories_models.GardenCategories, error) {
	ctx, ctxCancel := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)
	defer ctxCancel()

	collection := mongo_connection.MongoCollection(client, "gardenCategories")

	var results []garden_categories_models.GardenCategories

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println("Error finding all gardenCategories:", err)
	}

	cursorErr := cursor.All(context.TODO(), &results)

	if cursorErr != nil {
		log.Println(cursorErr)
	}

	return results, cursorErr
}