package flowers_models

type Flowers struct {
	ID       string `bson:"_id,omitempty" json:"_id"`
	Name     string `bson:"name" json:"name"`
	ImageURL string `bson:"imageURL" json:"imageURL"`
	Price    int    `bson:"price" json:"price"`
	IsActive bool   `bson:"isActive" json:"isActive"`
	IsSecret bool   `bson:"isSecret" json:"isSecret"`
}
