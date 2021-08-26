package gardens_models

type Rules struct {
	ID             string `bson:"_id" json:"_id"`
	Name           string `bson:"name" json:"name"`
	Description    string `bson:"description" json:"description"`
	UserFireBaseId string `bson:"userFireBaseId" json:"userFireBaseId"`
}
