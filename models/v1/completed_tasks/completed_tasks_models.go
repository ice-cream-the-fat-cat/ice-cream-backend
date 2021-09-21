package completed_tasks_models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompletedTasks struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	FireBaseUserId string `bson:"fireBaseUserId" json:"fireBaseUserId"`
  RuleId primitive.ObjectID `bson:"ruleId" json:"ruleId"`
  Date time.Time `bson:"date" json:"date"`
	RewardTypeId primitive.ObjectID `bson:"rewardTypeId" json:"rewardTypeId"`
	CreatedDate time.Time `bson:"createdDate" json:"createdDate"`
	LastUpdate time.Time `bson:"lastUpdate" json:"lastUpdate"`
}

func CompletedTaskValidation(completedTask CompletedTasks) bool {
	return completedTask.FireBaseUserId != "" && completedTask.RuleId != primitive.NilObjectID && !completedTask.Date.IsZero()
}