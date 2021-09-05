package users_models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	FireBaseUserId    string             `bson:"fireBaseUserId" json:"fireBaseUserId"`
	NumCoins          int                `bson:"numCoins" json:"numCoins"`
	FlowerCollections []primitive.ObjectID           `bson:"flowerCollections" json:"flowerCollections"`
	CreatedDate time.Time `bson:"createdDate" json:"createdDate"`
	LastUpdate time.Time `bson:"lastUpdate" json:"lastUpdate"`
}
