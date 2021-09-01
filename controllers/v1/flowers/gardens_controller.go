package flowers_controllers

import (
	"context"
	"log"

	mongo_connection "github.com/ice-cream-backend/database"
	flowers_models "github.com/ice-cream-backend/models/v1/flowers"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFlowerList() []flowers_models.Flowers {
	ctx := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)

	collection := mongo_connection.MongoCollection(client, "flowers")

	var results []flowers_models.Flowers

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
	}

	cursorErr := cursor.All(context.TODO(), &results)

	if cursorErr != nil {
		log.Println(cursorErr)
	}

	return results
}
