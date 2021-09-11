package mongo_connection

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongoErrorHandling(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func ContextForMongo() (context.Context, context.CancelFunc) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx, ctxCancel
}

func MongoConnection(ctx context.Context) mongo.Client {
	mongoURI := os.Getenv("MONGO_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	mongoErrorHandling(err)

	err = client.Connect(ctx)
	mongoErrorHandling(err)

	return *client
}

func MongoCollection(client mongo.Client, collectionName string) mongo.Collection {
	database := os.Getenv("MONGO_DB")
	collection := client.Database(database).Collection(collectionName)
	return *collection
}