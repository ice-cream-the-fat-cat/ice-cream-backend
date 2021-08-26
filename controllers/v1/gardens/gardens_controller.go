package gardens_controllers

import (
	"log"

	completed_tasks_controllers "github.com/ice-cream-backend/controllers/v1/completed_tasks"
	rules_controllers "github.com/ice-cream-backend/controllers/v1/rules"
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

func GetGardensByGardenId(createGardenId interface{}) gardens_models.Gardens {
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

func GetPopulatedGardenByGardenId(gardenId interface{}) gardens_models.GardensFullyPopulated {
	garden := GetGardensByGardenId(gardenId)
	rules := rules_controllers.GetRulesByGardenId(gardenId)

	var ruleIds []interface{}
	for _, rule := range rules {
		ruleIds = append(ruleIds, rule.ID)
	}

	completedTasks := completed_tasks_controllers.GetCompletedTasksByRuleIds(ruleIds)

	var populatedGarden gardens_models.GardensFullyPopulated
	populatedGarden.Garden = garden
	populatedGarden.Rules = rules
	populatedGarden.CompletedTasks = completedTasks

	return populatedGarden
}