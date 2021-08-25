package rules_controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	rules_models "github.com/ice-cream-backend/rules/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateRules(rulesPost rules_models.Rules) (*mongo.InsertOneResult, error) {
	log.Println("came into create rules controller with post:", rulesPost)
	fmt.Printf("%+v\n", rulesPost)

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

	// doc := bson.D{
	// 	primitive.E{Key: "name", Value: rulesPost.Name },
	// 	primitive.E{Key: "description", Value: rulesPost.Description },
	// 	primitive.E{Key: "gardenId", Value: rulesPost.GardenId },
	// 	primitive.E{Key: "isRemoved", Value: rulesPost.IsRemoved },
	// }

	res, insertErr := collection.InsertOne(ctx, rulesPost)

	if insertErr != nil {
		log.Println("Error creating new rule:", insertErr)
	}

	return res, insertErr
}