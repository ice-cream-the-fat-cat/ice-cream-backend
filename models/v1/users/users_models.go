package users_models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	FireBaseUserId    string             `bson:"fireBaseUserId" json:"fireBaseUserId"`
	NumCoins          int                `bson:"numCoins" json:"numCoins"`
	FlowerCollections []string           `bson:"flowerCollections" json:"flowerCollections"`
}
