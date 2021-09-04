package users_models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	//TODO : should check #79
	UserFireBaseId    string   `bson:"userFireBaseId" json:"userFireBaseId"`
	NumCoins          int      `bson:"numCoins" json:"numCoins"`
	FlowerCollections []string `bson:"flowerCollections" json:"flowerCollections"`
}
