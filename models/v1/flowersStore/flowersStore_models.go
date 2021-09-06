package flowersStore_models

type FlowersStore struct {
	FireBaseUserId string `bson:"fireBaseUserId" json:"fireBaseUserId"`
	FlowerID       string `bson:"_flowerId,omitempty" json:"_flowerId"`
	Price          int    `bson:"price" json:"price"`
}
