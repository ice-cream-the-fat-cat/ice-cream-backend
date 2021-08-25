package rules_models

type Rules struct {
	ID string `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
	Description string `bson:"description,omitempty"`
	GardenId string `bson:"gardenId,omitempty"`
	IsRemoved bool `bson:"isRemoved,omitempty"`
}

type RulesPost struct {
	Name string `json:"name"`
	Description int64 `json:"description"`
	GardenId string `json:"gardenId"`
	IsRemoved string `json:"isRemoved"`
}