package users_models

type Users struct {
	ID                string   `bson:"_id,omitempty" json:"_id"`
	UserFireBaseId    string   `bson:"userFireBaseId" json:"userFireBaseId"`
	NumCoins          int      `bson:"numCoins" json:"numCoins"`
	FlowerCollections []string `bson:"flowerCollections" json:"flowerCollections"`
}
