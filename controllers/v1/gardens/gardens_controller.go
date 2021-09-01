package gardens_controllers

import (
	"context"
	"log"
	"time"

	completed_tasks_controllers "github.com/ice-cream-backend/controllers/v1/completed_tasks"
	rules_controllers "github.com/ice-cream-backend/controllers/v1/rules"
	mongo_connection "github.com/ice-cream-backend/database"
	completed_tasks_models "github.com/ice-cream-backend/models/v1/completed_tasks"
	gardens_models "github.com/ice-cream-backend/models/v1/gardens"
	rules_models "github.com/ice-cream-backend/models/v1/rules"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateGardens(createdGardensPost gardens_models.Gardens) (*mongo.InsertOneResult, error) {
	ctx := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)

	collection := mongo_connection.MongoCollection(client, "gardens")

	createdGardensPost.CreatedDate = time.Now()
	createdGardensPost.LastUpdate = time.Now()

	res, insertErr := collection.InsertOne(ctx, createdGardensPost)

	if insertErr != nil {
		log.Println("Error creating new createGardens:", insertErr)
	}

	return res, insertErr
}

func GetGardensByGardenId(createGardenId interface{}) (gardens_models.Gardens, error) {
	ctx := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)

	collection := mongo_connection.MongoCollection(client, "gardens")

	var result gardens_models.Gardens
	err := collection.FindOne(ctx, bson.D{
		primitive.E{Key: "_id", Value: createGardenId},
	},
	).Decode(&result)
	
	if err != nil {
		log.Println("err in findOne:", err)
	}

	return result, err
}

func GetPopulatedGardenByGardenId(gardenId interface{}) (gardens_models.GardensFullyPopulated, error) {
	garden, err := GetGardensByGardenId(gardenId)

	var populatedGarden gardens_models.GardensFullyPopulated

	if err != nil {
		return populatedGarden, err
	}

	rules := rules_controllers.GetRulesByGardenId(gardenId)

	var ruleIds []interface{}
	for _, rule := range rules {
		ruleIds = append(ruleIds, rule.ID)
	}

	populatedGarden.Garden = garden

	if len(ruleIds) == 0 {
		populatedGarden.Rules = []rules_models.Rules{}
		populatedGarden.CompletedTasks = []completed_tasks_models.CompletedTasks{}
	} else{
		populatedGarden.Rules = rules
		completedTasks := completed_tasks_controllers.GetCompletedTasksByRuleIds(ruleIds)
		populatedGarden.CompletedTasks = completedTasks
	}

	return populatedGarden, nil
}

func GetGardensByUserId(userFireBaseId interface{}) []gardens_models.Gardens {
	ctx := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)

	collection := mongo_connection.MongoCollection(client, "gardens")

	var results []gardens_models.Gardens
	query := bson.D{
		primitive.E{Key: "userFireBaseId", Value: userFireBaseId},
	}
	cursor, err := collection.Find(ctx, query)
	if err != nil {
		log.Println(err)
	}

	cursorErr := cursor.All(context.TODO(), &results)

	if cursorErr != nil {
		log.Println(cursorErr)
	}

	return results
}

func UpdateGardenByGardenId(gardenId interface{}, garden gardens_models.Gardens) (*mongo.UpdateResult, error) {
	ctx := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)

	collection := mongo_connection.MongoCollection(client, "gardens")

	updatedGarden := bson.M{
		"$set": bson.M{
			"name": garden.Name,
			"description": garden.Description,
			"lastUpdate": time.Now(),
		},
	}

	result, updateErr := collection.UpdateByID(ctx, gardenId, updatedGarden)

	return result, updateErr
}

func DeleteGardenByGardenId(gardenId interface{}) (*mongo.DeleteResult, error) {
	ctx := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)

	collection := mongo_connection.MongoCollection(client, "gardens")

	query := bson.D{
		primitive.E{Key: "_id", Value: gardenId},
	}
	gardenRes, gardenErr := collection.DeleteOne(context.TODO(), query)

	if gardenErr != nil {
		log.Println(gardenErr)
	}

	_, rulesErr := rules_controllers.DeleteRulesByGardenId(gardenId)

	if rulesErr != nil {
		log.Println("Error deleting rules for ID", gardenId, rulesErr)
	}

	return gardenRes, gardenErr
}