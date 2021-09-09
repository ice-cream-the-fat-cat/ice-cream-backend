package gardens_models

import (
	"time"

	completed_tasks_models "github.com/ice-cream-backend/models/v1/completed_tasks"
	garden_categories_models "github.com/ice-cream-backend/models/v1/garden_categories"
	rules_models "github.com/ice-cream-backend/models/v1/rules"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gardens struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name           string `bson:"name" json:"name"`
	Description    string `bson:"description" json:"description"`
	FireBaseUserId string `bson:"fireBaseUserId" json:"fireBaseUserId"`
	GardenCategoryId primitive.ObjectID `bson:"gardenCategoryId" json:"gardenCategoryId"`
	GardenCategory garden_categories_models.GardenCategories `json:"gardenCategory"`
	CreatedDate time.Time `bson:"createdDate" json:"createdDate"`
	LastUpdate time.Time `bson:"lastUpdate" json:"lastUpdate"`
}

type GardenForMongo struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name           string `bson:"name" json:"name"`
	Description    string `bson:"description" json:"description"`
	FireBaseUserId string `bson:"fireBaseUserId" json:"fireBaseUserId"`
	GardenCategoryId primitive.ObjectID `bson:"gardenCategoryId" json:"gardenCategoryId"`
	CreatedDate time.Time `bson:"createdDate" json:"createdDate"`
	LastUpdate time.Time `bson:"lastUpdate" json:"lastUpdate"`
}

type GardensFullyPopulated struct {
	Garden         Gardens                                 `json:"garden"`
	Rules          []rules_models.Rules                    `json:"rules"`
	CompletedTasks []completed_tasks_models.CompletedTasks `json:"completedTasks"`
}
