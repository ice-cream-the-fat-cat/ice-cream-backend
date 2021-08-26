package completedTasks_models

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
	CreatedDate time.Time `bson:"createdAt" json:"createdAt"`
	LastUpdate time.Time `bson:"lastUpdate" json:"lastUpdate"`
}