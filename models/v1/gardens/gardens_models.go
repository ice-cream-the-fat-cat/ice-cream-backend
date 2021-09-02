package gardens_models

import (
	"time"

	completed_tasks_models "github.com/ice-cream-backend/models/v1/completed_tasks"
	rules_models "github.com/ice-cream-backend/models/v1/rules"
)

type Gardens struct {
	ID             string `bson:"_id,omitempty" json:"_id"`
	Name           string `bson:"name" json:"name"`
	Description    string `bson:"description" json:"description"`
	UserFireBaseId string `bson:"userFireBaseId" json:"userFireBaseId"`
	CreatedDate time.Time `bson:"createdDate" json:"createdDate"`
	LastUpdate time.Time `bson:"lastUpdate" json:"lastUpdate"`
}

type GardensFullyPopulated struct {
	Garden         Gardens                                 `json:"garden"`
	Rules          []rules_models.Rules                    `json:"rules"`
	CompletedTasks []completed_tasks_models.CompletedTasks `json:"completedTasks"`
}
