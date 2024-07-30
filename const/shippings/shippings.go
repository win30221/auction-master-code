package shippings

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// 出貨方式類型
	AddressType     uint8 = 1 // 地址寄出
	SevenElevenType uint8 = 2 // 7-11寄出
	FamilyType      uint8 = 3 // 全家寄出

	// 出貨狀態
	SubmitAppraisalStatus uint8 = 1 // 已提交出貨
	ProcessingStatus      uint8 = 2 // 理貨中
	ShippedStatus         uint8 = 3 // 已寄出

)

type Shipping struct {
	ID            *primitive.ObjectID `json:"id" bson:"_id"`
	Type          *uint8              `form:"type" json:"type" bson:"type" sendForm:"type"`
	ItemIDs       *[]int              `form:"itemIDs" json:"itemIDs" bson:"itemIDs" sendForm:"itemIDs"`
	Address       *string             `form:"address" json:"address,omitempty" bson:"address,omitempty" sendForm:"address"`
	StoreNumber   *string             `form:"storeNumber" json:"storeNumber,omitempty" bson:"storeNumber,omitempty" sendForm:"storeNumber"`
	StoreName     *string             `form:"storeName" json:"storeName,omitempty" bson:"storeName,omitempty" sendForm:"storeName"`
	RecipientName *string             `form:"recipientName" json:"recipientName" bson:"recipientName" sendForm:"recipientName"`
	Phone         *string             `form:"phone" json:"phone" bson:"phone" sendForm:"phone"`
	Status        *uint8              `form:"status" json:"status" bson:"status" sendForm:"status"`
	CreatedAt     *time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt     *time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type GetShippingsRes struct {
	Shippings []Shipping `json:"shippings"`
	Count     int        `json:"count"`
}
