package flowersStore_models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FlowersStore struct {
	FireBaseUserId string `bson:"fireBaseUserId" json:"fireBaseUserId"`
	FlowerID       primitive.ObjectID `bson:"flowerId" json:"flowerId"`
	Price          int    `bson:"price" json:"price"`
}

func FlowerStoreValidation(flowerStore FlowersStore) bool {
	return flowerStore.FireBaseUserId != "" && flowerStore.FlowerID != primitive.NilObjectID && flowerStore.Price != 0
}