package rules_controllers

import (
	"context"
	"log"
	"time"

	mongo_connection "github.com/ice-cream-backend/database"
	rules_models "github.com/ice-cream-backend/models/v1/rules"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateRule(rulesPost rules_models.Rules) (*mongo.InsertOneResult, error) {
	ctx := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)

	collection := mongo_connection.MongoCollection(client, "rules")

	rulesPost.CreatedDate = time.Now()
	rulesPost.LastUpdate = time.Now()

	res, insertErr := collection.InsertOne(ctx, rulesPost)

	if insertErr != nil {
		log.Println("Error creating new rule:", insertErr)
	}

	return res, insertErr
}

func CreateRules(multipleRulesPost []rules_models.Rules) (*mongo.InsertManyResult, error) {
	ctx := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)

	collection := mongo_connection.MongoCollection(client, "rules")

	var populatedRules []interface{}
	for _, rule := range multipleRulesPost {
		rule.CreatedDate = time.Now()
		rule.LastUpdate = time.Now()
		populatedRules = append(populatedRules, rule)
	}

	opts := options.InsertMany().SetOrdered(false)
	res, insertErr := collection.InsertMany(ctx, populatedRules, opts)

	if insertErr != nil {
		log.Println("Error creating new rule:", insertErr)
	}

	return res, insertErr
}

func GetRulesByRuleId(ruleId interface{}) rules_models.Rules {
	ctx := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)

	collection := mongo_connection.MongoCollection(client, "rules")

	var result rules_models.Rules
	collection.FindOne(ctx, bson.D{
		primitive.E{Key:"_id", Value: ruleId},
		},
	).Decode(&result)

	return result
}

func GetRulesByRuleIds(ruleIds []interface{}) []rules_models.Rules {
	ctx := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)

	collection := mongo_connection.MongoCollection(client, "rules")

	var results []rules_models.Rules
	query := bson.M{"_id": bson.M{"$in": ruleIds}}
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

func GetRulesByGardenId(gardenId interface{}) []rules_models.Rules {
	ctx := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)

	collection := mongo_connection.MongoCollection(client, "rules")

	var results []rules_models.Rules
	query := bson.D{
		primitive.E{Key:"gardenId", Value: gardenId},
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

func UpdateRuleByRuleId(ruleId interface{}, rule rules_models.Rules) (*mongo.UpdateResult, error) {
	ctx := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)

	collection := mongo_connection.MongoCollection(client, "rules")

	updatedRule := bson.M{
    "$set": bson.M{
      "name": rule.Name,
      "description": rule.Description,
			"isRemoved": rule.IsRemoved,
			"lastUpdate": time.Now(),
    },
  }

	result, updateErr := collection.UpdateByID(ctx, ruleId, updatedRule)

	return result, updateErr
}