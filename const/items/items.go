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

type Item struct {
	ID                 *uint64    `form:"id" json:"id"`
	ConsignorID        *uint64    `form:"consignorID" json:"consignorID"`
	Name               *string    `form:"name" json:"name"`
	Description        *string    `form:"description" json:"description"`
	Weight             *int       `form:"weight" json:"weight"`
	ShippingFeeToJapan *int       `form:"shippingFeeToJapan" json:"shippingFeeToJapan"`
	MinEstimatedPrice  *int       `form:"minEstimatedPrice" json:"minEstimatedPrice"`
	MaxEstimatedPrice  *int       `form:"maxEstimatedPrice" json:"maxEstimatedPrice"`
	Cost               *int       `form:"cost" json:"cost"`
	ReservePrice       *int       `form:"reservePrice" json:"reservePrice"`
	IncludeUserCost    *uint8     `form:"includeUserCost" json:"includeUserCost"`
	Status             *uint8     `form:"status" json:"status"`
	CreatedAt          *time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt          *time.Time `form:"updatedAt" json:"updatedAt"`
}
