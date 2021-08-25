package rules_models

type Rules struct {
	ID string `bson:"_id" json:"_id"`
	Name string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	GardenId string `bson:"gardenId" json:"gardenId"`
	IsRemoved bool `bson:"isRemoved" json:"isRemoved"`
}
