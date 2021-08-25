package rules_models

type Rules struct {
	ID string `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
	Description string `bson:"description,omitempty"`
	GardenId string `bson:"gardenId,omitempty"`
	IsRemoved bool `bson:"isRemoved,omitempty"`
}