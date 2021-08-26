package rules_controllers

import (
	"context"
	"log"
	"os"
	"time"

	rules_models "github.com/ice-cream-backend/models/v1/rules"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateRules(rulesPost rules_models.Rules) (*mongo.InsertOneResult, error) {
	// TODO: Move MongoDB connection to utils
	mongoURI := os.Getenv("MONGO_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("icecream-dev").Collection("rules")

	rulesPost.CreatedDate = time.Now()
	rulesPost.LastUpdate = time.Now()

	res, insertErr := collection.InsertOne(ctx, rulesPost)

	if insertErr != nil {
		log.Println("Error creating new rule:", insertErr)
	}

	return res, insertErr
}

func GetRules(ruleId interface{}) rules_models.Rules {
	mongoURI := os.Getenv("MONGO_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	database := os.Getenv("MONGO_DB")
	collection := client.Database(database).Collection("rules")

	var result rules_models.Rules
	collection.FindOne(ctx, bson.D{
		primitive.E{Key:"_id", Value: ruleId},
		},
	).Decode(&result)

	return result
}