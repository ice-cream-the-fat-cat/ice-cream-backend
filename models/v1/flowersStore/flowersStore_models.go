package flowersStore_models

type FlowersStore struct {
	FireBaseUserId string `bson:"fireBaseUserId" json:"fireBaseUserId"`
	FlowerID       string `bson:"flowerId" json:"flowerId"`
	Price          int    `bson:"price" json:"price"`
}
