package items

import "time"

var (
	// 商品類型
	AppraisableAuctionItemType    uint8 = 1 // 可估價競標物品
	NonAppraisableAuctionItemType uint8 = 2 // 不可估價競標物品
	FixedPriceItemType            uint8 = 3 // 定價物品

	// 商品狀態
	InitStatus             uint8 = 1 // 初始化
	SubmitAppraisalStatus  uint8 = 2 // 已提交估價
	AppraisalFailureStatus uint8 = 3 // 估價失敗
	AppraisedStatus        uint8 = 4 // 已估價

	ConsignmentApprovedStatus uint8 = 11 // 同意託售
	ConsignmentCanceledStatus uint8 = 12 // 取消託售

	WarehouseArrivalStatus       uint8 = 21 // 物品到達倉庫
	WarehouseReturnPendingStatus uint8 = 22 // 到倉庫，準備退回中
	DetailsFullyCompletedStatus  uint8 = 23 // 上架資料齊全
	ReadyStatus                  uint8 = 24 // 準備上架
	BiddingStatus                uint8 = 25 // 競標中

	SoldStatus               uint8 = 31 // 已售出
	ReturnedStatus           uint8 = 32 // 退回
	CompanyRepurchasedStatus uint8 = 33 // 被公司買回
	CompanyReclaimedStatus   uint8 = 34 // 被公司收回
)

type ItemDetails struct {
	ID                *uint64     `json:"id"`
	ConsignorID       *uint64     `json:"consignorID"`
	Type              *uint8      `json:"type"`
	Name              *string     `json:"name"`
	Description       *string     `json:"description"`
	Photos            []ItemPhoto `json:"photos"`
	Space             *uint8      `json:"space"`
	MinEstimatedPrice *int        `json:"minEstimatedPrice"`
	MaxEstimatedPrice *int        `json:"maxEstimatedPrice"`
	SellerID          *uint64     `json:"sellerID"`
	ReservePrice      *int        `json:"reservePrice"`
	ExpireAt          *time.Time  `json:"expireAt"`
	Status            *uint8      `json:"status"`
	CreatedAt         *time.Time  `json:"createdAt"`
	UpdatedAt         *time.Time  `json:"updatedAt"`
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
	SellerID          *uint64    `form:"sellerID" json:"sellerID"`
	ReservePrice      *int       `form:"reservePrice" json:"reservePrice"`
	ExpireAt          *time.Time `form:"expireAt" json:"expireAt"`
	Status            *uint8     `form:"status" json:"status"`
	CreatedAt         *time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt         *time.Time `form:"updatedAt" json:"updatedAt"`
}

type ItemPhoto struct {
	ItemID    *uint64    `form:"itemID" json:"itemID,omitempty"`
	Sorted    *uint8     `form:"sorted" json:"sorted"`
	Photo     *string    `form:"photo" json:"photo"`
	CreatedAt *time.Time `form:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `form:"updatedAt" json:"updatedAt,omitempty"`
}

type GetItemsAndDetailsRes struct {
	Items []ItemDetails `json:"items"`
	Count int           `json:"count"`
}
