package items

import "time"

var (
	// 商品類型
	AppraisableAuctionItemType    uint8 = 1 // 可估價競標物品
	NonAppraisableAuctionItemType uint8 = 2 // 不可估價競標物品
	FixedPriceItemType            uint8 = 3 // 定價物品

	// 商品狀態
	SubmitAppraisalStatus  uint8 = 1 // 已提交估價
	AppraisalFailureStatus uint8 = 2 // 估價失敗
	AppraisedStatus        uint8 = 3 // 已估價

	ConsignmentApprovedStatus uint8 = 11 // 同意託售
	ConsignmentCanceledStatus uint8 = 12 // 取消託售

	WarehouseReturnPendingStatus uint8 = 21 // 準備退回(等待寄售人聯絡退貨事宜)
	WarehouseReturningStatus     uint8 = 22 // 倉庫人員退貨準備中
	WarehouseArrivalStatus       uint8 = 23 // 倉庫人員確認物品到達倉庫
	DetailsFullyCompletedStatus  uint8 = 24 // 客服確認物品可出售
	ReadyStatus                  uint8 = 25 // 寄售人確認可準備上架
	BiddingStatus                uint8 = 26 // 競標中

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
	Sorted    *uint8     `form:"sorted" validate:"required,gt=0" json:"sorted"`
	Photo     *string    `form:"photo" json:"photo"`
	CreatedAt *time.Time `form:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `form:"updatedAt" json:"updatedAt,omitempty"`
}

type ReorderItemPhotoReq struct {
	OriginalSorted uint8 `form:"originalSorted" validate:"required"`
	NewSorted      uint8 `form:"newSorted" validate:"required"`
}

type GetItemsAndDetailsRes struct {
	Items []ItemDetails `json:"items"`
	Count int           `json:"count"`
}
