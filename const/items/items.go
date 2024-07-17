package items

import "time"

var (
	// 商品類型
	AppraisableAuctionItemType    uint8 = 1 // 可估價競標物品
	NonAppraisableAuctionItemType uint8 = 2 // 不可估價競標物品
	FixedPriceItemType            uint8 = 3 // 定價物品
	CompanyDirectPurchaseType     uint8 = 4 // 公司直購物品

	// 商品狀態
	SubmitAppraisalStatus  uint8 = 1 // 已提交估價
	AppraisalFailureStatus uint8 = 2 // 估價失敗
	AppraisedStatus        uint8 = 3 // 已估價

	ConsignmentApprovedStatus                   uint8 = 11 // 同意託售
	ConsignmentCanceledStatus                   uint8 = 12 // 取消託售
	ConsignorChoosesCompanyDirectPurchaseStatus uint8 = 13 // 寄售人選擇公司直購
	ConsignorShippedItem                        uint8 = 14 // 寄售人已寄出

	WarehouseArrivalStatus            uint8 = 21 // 已到貨
	WarehouseReturnPendingStatus      uint8 = 22 // 準備退回(等待寄售人聯絡退貨事宜)
	WarehouseReturningStatus          uint8 = 23 // 退貨作業中
	WarehousePersonnelConfirmedStatus uint8 = 24 // 倉管已確認
	AppraiserConfirmedStatus          uint8 = 25 // 鑑價師已確認
	ConsignorConfirmedStatus          uint8 = 26 // 寄售人已確認可準備上架
	BiddingStatus                     uint8 = 27 // 競標中

	SoldStatus                  uint8 = 31 // 已售出
	CompanyDirectPurchaseStatus uint8 = 32 // 公司直購
	ReturnedStatus              uint8 = 33 // 退回
	CompanyRepurchasedStatus    uint8 = 34 // 公司買回
	CompanyReclaimedStatus      uint8 = 35 // 公司沒收

)

type ItemDetails struct {
	ID                  *uint64          `json:"id"`
	ConsignorID         *uint64          `json:"consignorID"`
	Type                *uint8           `json:"type"`
	Name                *string          `json:"name"`
	Description         *string          `json:"description"`
	Photos              []ItemPhoto      `json:"photos"`
	PastStatuses        []ItemPastStatus `json:"pastStatuses"`
	DirectPurchasePrice *int             `json:"directPurchasePrice"`
	MinEstimatedPrice   *int             `json:"minEstimatedPrice"`
	MaxEstimatedPrice   *int             `json:"maxEstimatedPrice"`
	ReservePrice        *int             `json:"reservePrice"`
	ExpireAt            *time.Time       `json:"expireAt"`
	WarehouseID         *string          `json:"warehouseID"`
	Space               *uint8           `json:"space"`
	GrossWeight         *int             `json:"grossWeight"`
	VolumetricWeight    *int             `json:"volumetricWeight"`
	Status              *uint8           `json:"status"`
	CreatedAt           *time.Time       `json:"createdAt"`
	UpdatedAt           *time.Time       `json:"updatedAt"`
}

type Item struct {
	ID                  *uint64    `form:"id" json:"id"`
	ConsignorID         *uint64    `form:"consignorID" json:"consignorID"`
	Type                *uint8     `form:"type" json:"type"`                               // 物品類型
	Name                *string    `form:"name" json:"name"`                               // 物品名稱
	Description         *string    `form:"description" json:"description"`                 // 物品描述
	DirectPurchasePrice *int       `form:"directPurchasePrice" json:"directPurchasePrice"` // 直購金額
	MinEstimatedPrice   *int       `form:"minEstimatedPrice" json:"minEstimatedPrice"`     // 最低估價金額
	MaxEstimatedPrice   *int       `form:"maxEstimatedPrice" json:"maxEstimatedPrice"`     // 最高估價金額
	ReservePrice        *int       `form:"reservePrice" json:"reservePrice"`               // 期望金額(用於最後盯標)
	ExpireAt            *time.Time `form:"expireAt" json:"expireAt"`                       // 過期時間
	WarehouseID         *string    `form:"warehouseID" json:"warehouseID"`                 // 存放於倉庫的編號
	Space               *uint8     `form:"space" json:"space"`                             // 物品占用空間
	GrossWeight         *int       `form:"grossWeight" json:"grossWeight"`                 // 實際重量
	VolumetricWeight    *int       `form:"volumetricWeight" json:"volumetricWeight"`       // 體積重量
	Status              *uint8     `form:"status" json:"status"`
	CreatedAt           *time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt           *time.Time `form:"updatedAt" json:"updatedAt"`
}

type ItemPhoto struct {
	ItemID    *uint64    `form:"itemID" json:"itemID,omitempty"`
	Sorted    *uint8     `form:"sorted" validate:"required,gt=0" json:"sorted"`
	Photo     *string    `form:"photo" json:"photo"`
	CreatedAt *time.Time `form:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `form:"updatedAt" json:"updatedAt,omitempty"`
}

type ItemPastStatus struct {
	ItemID    *uint64    `form:"itemID" json:"itemID,omitempty"`
	Status    *uint8     `form:"status" json:"status"`
	CreatedAt *time.Time `form:"createdAt" json:"createdAt,omitempty"`
}

type ReorderItemPhotoReq struct {
	OriginalSorted uint8 `form:"originalSorted" validate:"required"`
	NewSorted      uint8 `form:"newSorted" validate:"required"`
}

type GetItemsAndDetailsRes struct {
	Items []ItemDetails `json:"items"`
	Count int           `json:"count"`
}
