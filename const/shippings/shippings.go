package shippings

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// 出貨目的類型
	YahooDispatchActionType uint8 = 1 // 日拍出貨
	ReturnActionType        uint8 = 2 // 退貨

	// 出貨方式類型
	AddressShipmentType     uint8 = 1 // 地址寄出
	SevenElevenShipmentType uint8 = 2 // 7-11寄出
	FamilyShipmentType      uint8 = 3 // 全家寄出

	// 出貨狀態
	AwaitingConsignorPayFeeStatus uint8 = 1  // 等待寄售人付費
	SubmitAppraisalStatus         uint8 = 2  // 已提交出貨
	ProcessingStatus              uint8 = 3  // 理貨中
	ShippedStatus                 uint8 = 4  // 已寄出
	CanceledStatus                uint8 = 99 // 取消

)

type Shipping struct {
	ID                       *primitive.ObjectID `json:"id" bson:"_id"`
	ActionType               *uint8              `form:"actionType" json:"actionType" bson:"actionType"`
	ShipmentType             *uint8              `form:"shipmentType" json:"shipmentType" bson:"shipmentType"`
	ItemIDs                  *[]uint64           `form:"itemIDs" json:"itemIDs" bson:"itemIDs"`
	AuctionItemIDs           *[]uint64           `form:"auctionItemIDs" json:"auctionItemIDs,omitempty" bson:"auctionItemIDs,omitempty"`
	Address                  *string             `form:"address" json:"address,omitempty" bson:"address,omitempty"`
	StoreNumber              *string             `form:"storeNumber" json:"storeNumber,omitempty" bson:"storeNumber,omitempty"`
	StoreName                *string             `form:"storeName" json:"storeName,omitempty" bson:"storeName,omitempty"`
	RecipientName            *string             `form:"recipientName" json:"recipientName" bson:"recipientName"`
	Phone                    *string             `form:"phone" json:"phone" bson:"phone"`
	ShipmentTrackingNumber   *string             `form:"shipmentTrackingNumber" json:"shipmentTrackingNumber" bson:"shipmentTrackingNumber"`
	ShippingCostsWithinJapan *int                `form:"shippingCostsWithinJapan" json:"shippingCostsWithinJapan,omitempty" bson:"shippingCostsWithinJapan,omitempty"`
	Status                   *uint8              `form:"status" json:"status" bson:"status"`
	CreatedAt                *time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt                *time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type GetShippingsReq struct {
	Status  []int     `form:"status"` // mongo 查詢沒有 uint8 類型
	StartAt time.Time `form:"startAt"`
	EndAt   time.Time `form:"endAt"`
	Sort    []string  `form:"sort"`
	Order   []string  `form:"order"`
	Limit   int64     `form:"limit"`
	Offset  int64     `form:"offset"`
}

type GetShippingsRes struct {
	Shippings []Shipping `json:"shippings"`
	Count     int64      `json:"count"`
}
