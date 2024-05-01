package items

import "time"

var (
	// 競標商品狀態
	InitStatus      uint8 = 1 // 初始化
	AppraisedStatus uint8 = 2 // 已估價
	CanceledStatus  uint8 = 3 // 取消

	InWarehouseStatus           uint8 = 11 // 到倉庫(準備資料中)
	DetailsFullyCompletedStatus uint8 = 12 // 上架資料齊全

	SoldStatus     uint8 = 21 // 已售出
	ReturnedStatus uint8 = 22 // 退回

)

type ItemDetails struct {
	ID                *uint64    `json:"id"`
	ConsignorID       *uint64    `json:"consignorID"`
	Type              *uint8     `json:"type"`
	Name              *string    `json:"name"`
	Description       *string    `json:"description"`
	Photo             *string    `json:"photo"`
	Space             *uint8     `json:"space"`
	MinEstimatedPrice *int       `json:"minEstimatedPrice"`
	MaxEstimatedPrice *int       `json:"maxEstimatedPrice"`
	ReservePrice      *int       `json:"reservePrice"`
	ExpireAt          *time.Time `json:"expireAt"`
	Status            *uint8     `json:"status"`
	CreatedAt         *time.Time `json:"createdAt"`
	UpdatedAt         *time.Time `json:"updatedAt"`
}

type Item struct {
	ID                *uint64    `form:"id" json:"id"`
	ConsignorID       *uint64    `form:"consignorID" json:"consignorID"`
	Type              *uint8     `form:"type" json:"type"`
	Name              *string    `form:"name" json:"name"`
	Description       *string    `form:"description" json:"description"`
	Space             *uint8     `form:"space" json:"space"`
	MinEstimatedPrice *int       `form:"minEstimatedPrice" json:"minEstimatedPrice"`
	MaxEstimatedPrice *int       `form:"maxEstimatedPrice" json:"maxEstimatedPrice"`
	ReservePrice      *int       `form:"reservePrice" json:"reservePrice"`
	ExpireAt          *time.Time `form:"expireAt" json:"expireAt"`
	Status            *uint8     `form:"status" json:"status"`
	CreatedAt         *time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt         *time.Time `form:"updatedAt" json:"updatedAt"`
}

type ItemPhoto struct {
	ItemID    *uint64    `form:"itemID" json:"itemID"`
	Index     *uint8     `form:"index" json:"index"`
	URL       *string    `form:"url" json:"url"`
	CreatedAt *time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time `form:"updatedAt" json:"updatedAt"`
}

type GetItemDetailsRes struct {
	Items []ItemDetails `json:"items"`
	Count int           `json:"count"`
}
