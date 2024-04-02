package workerhubs

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// Worker 類型
	WorkerTypeSeller  = "Seller"
	WorkerTypeWatcher = "Watcher"
)

type Worker struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id" form:"id"`
	Type          string             `bson:"type" json:"type" form:"type"`
	URL           string             `bson:"url" json:"url" form:"url"`
	Status        uint8              `bson:"status" json:"status" form:"status"`
	Location      string             `bson:"location" json:"location" form:"location"`
	Name          string             `bson:"name" json:"name" form:"name"`
	Phone         string             `bson:"phone,omitempty" json:"phone,omitempty" form:"phone"`
	PostalCode    string             `bson:"postalCode,omitempty" json:"postalCode,omitempty" form:"postalCode"`
	Birthday      time.Time          `bson:"birthday,omitempty" json:"birthday,omitempty" form:"birthday"`
	Email         string             `bson:"email,omitempty" json:"email,omitempty" form:"email"`
	SimCardNumber string             `bson:"simCardNumber,omitempty" json:"simCardNumber,omitempty" form:"simCardNumber"`
}
