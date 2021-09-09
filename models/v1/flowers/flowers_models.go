package flowers_models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Flowers struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name     string `bson:"name" json:"name"`
	ImageURL string `bson:"imageURL" json:"imageURL"`
	Price    int    `bson:"price" json:"price"`
	IsActive bool   `bson:"isActive" json:"isActive"`
	IsSecret bool   `bson:"isSecret" json:"isSecret"`
}
