package shippings

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shippings struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Type          string             `form:"type" json:"type" bson:"type" sendForm:"type"`
	ItemIDs       []int              `form:"itemIDs" json:"itemIDs" bson:"itemIDs" sendForm:"itemIDs"`
	Address       string             `form:"address" json:"address,omitempty" bson:"address,omitempty" sendForm:"address"`
	StoreNumber   string             `form:"storeNumber" json:"storeNumber,omitempty" bson:"storeNumber,omitempty" sendForm:"storeNumber"`
	StoreName     string             `form:"storeName" json:"storeName,omitempty" bson:"storeName,omitempty" sendForm:"storeName"`
	RecipientName string             `form:"recipientName" json:"recipientName" bson:"recipientName" sendForm:"recipientName"`
	Phone         string             `form:"phone" json:"phone" bson:"phone" sendForm:"phone"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type GetShippingsRes struct {
	Shippings []Shippings `json:"shippings"`
	Count     int         `json:"count"`
}
