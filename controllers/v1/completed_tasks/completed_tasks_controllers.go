package completed_tasks_controllers

import (
	"context"
	"log"
	"time"

	mongo_connection "github.com/ice-cream-backend/database"
	completed_tasks_models "github.com/ice-cream-backend/models/v1/completed_tasks"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateCompletedTask(completedTasksPost completed_tasks_models.CompletedTasks) (*mongo.InsertOneResult, error) {
	ctx, ctxCancel := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)
	defer ctxCancel()

	collection := mongo_connection.MongoCollection(client, "completedTasks")

	completedTasksPost.CreatedDate = time.Now()
	completedTasksPost.LastUpdate = time.Now()

	res, insertErr := collection.InsertOne(ctx, completedTasksPost)

	if insertErr != nil {
		log.Println("Error creating new completedTasks:", insertErr)
	}

	return res, insertErr
}

func GetCompletedTasksByCompletedTaskId(completedTaskId interface{}) completed_tasks_models.CompletedTasks {
	ctx, ctxCancel := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)
	defer ctxCancel()

	collection := mongo_connection.MongoCollection(client, "completedTasks")

	var result completed_tasks_models.CompletedTasks
	collection.FindOne(ctx, bson.D{
		primitive.E{Key:"_id", Value: completedTaskId},
		},
	).Decode(&result)

	return result
}

func GetCompletedTasksByRuleIds(ruleIds []interface{}) []completed_tasks_models.CompletedTasks {
	ctx, ctxCancel := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)
	defer ctxCancel()

	collection := mongo_connection.MongoCollection(client, "completedTasks")

	var results []completed_tasks_models.CompletedTasks
	query := bson.M{"ruleId": bson.M{"$in": ruleIds}}
	opts := options.Find().SetSort(bson.D{
		primitive.E{Key:"date", Value: 1},
	})
	cursor, err := collection.Find(ctx, query, opts)
	if err != nil {
		log.Println(err)
	}

	cursorErr := cursor.All(context.TODO(), &results)

	if cursorErr != nil {
		log.Println(cursorErr)
	}

	return results
}

func GetCompletedTasksByRuleIdWithDate(ruleIds []interface{}, date time.Time) []completed_tasks_models.CompletedTasks {
	ctx, ctxCancel := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)
	defer ctxCancel()

	collection := mongo_connection.MongoCollection(client, "completedTasks")

	var results []completed_tasks_models.CompletedTasks
	query := bson.M{
		"ruleId": bson.M{"$in": ruleIds},
		"date": date,
	}
	opts := options.Find().SetSort(bson.D{
		primitive.E{Key:"date", Value: 1},
	})
	cursor, err := collection.Find(ctx, query, opts)
	if err != nil {
		log.Println(err)
	}

	cursorErr := cursor.All(context.TODO(), &results)

	if cursorErr != nil {
		log.Println(cursorErr)
	}

	return results
}

func GetCompletedTasksByRuleIdWithStartAndEndDate(ruleIds []interface{}, startDate time.Time, endDate time.Time) []completed_tasks_models.CompletedTasks {
	ctx, ctxCancel := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)
	defer ctxCancel()

	collection := mongo_connection.MongoCollection(client, "completedTasks")

	var results []completed_tasks_models.CompletedTasks
	query := bson.M{
		"ruleId": bson.M{"$in": ruleIds},
		"date": bson.M{"$gte": startDate , "$lte": endDate},
	}
	opts := options.Find().SetSort(bson.D{
		primitive.E{Key:"date", Value: 1},
	})
	cursor, err := collection.Find(ctx, query, opts)
	if err != nil {
		log.Println(err)
	}

	cursorErr := cursor.All(context.TODO(), &results)

	if cursorErr != nil {
		log.Println(cursorErr)
	}

	return results
}

func DeleteCompletedTaskByCompletedTaskId(completedTaskId interface{}) (*mongo.DeleteResult, error) {
	ctx, ctxCancel := mongo_connection.ContextForMongo()
	client := mongo_connection.MongoConnection(ctx)

	defer client.Disconnect(ctx)
	defer ctxCancel()

	collection := mongo_connection.MongoCollection(client, "completedTasks")

	query := bson.D{
		primitive.E{Key: "_id", Value: completedTaskId},
	}
	completedTaskRes, completedTaskErr := collection.DeleteOne(context.TODO(), query)

	if completedTaskErr != nil {
		log.Println("Error deleting completedTask:", completedTaskErr)
	}

	return completedTaskRes, completedTaskErr
}