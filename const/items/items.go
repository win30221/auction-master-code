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
	CompanyPurchasedStatus      uint8 = 34 // 公司買回
	CompanyReclaimedStatus      uint8 = 35 // 公司沒收

)

type Item struct {
	Id                       *uint64    `form:"id" json:"id"`
	ConsignorId              *uint64    `form:"consignorId" json:"consignorId"`
	Type                     *uint8     `form:"type" json:"type"`                                                   // 物品類型
	IsNew                    *bool      `form:"isNew" json:"isNew" `                                                // 是否全新品
	Name                     *string    `form:"name" json:"name"`                                                   // 物品名稱
	Description              *string    `form:"description" json:"description"`                                     // 物品描述
	DirectPurchasePrice      *int       `form:"directPurchasePrice" json:"directPurchasePrice"`                     // 直購金額
	MinEstimatedPrice        *int       `form:"minEstimatedPrice" json:"minEstimatedPrice"`                         // 最低估價金額
	MaxEstimatedPrice        *int       `form:"maxEstimatedPrice" json:"maxEstimatedPrice"`                         // 最高估價金額
	ReservePrice             *int       `form:"reservePrice" json:"reservePrice"`                                   // 期望金額(用於最後盯標)
	ExpireAt                 *time.Time `form:"expireAt" json:"expireAt"`                                           // 過期時間
	WarehouseId              *string    `form:"warehouseId" json:"warehouseId"`                                     // 存放於倉庫的編號
	Space                    *uint8     `form:"space" json:"space"`                                                 // 物品占用空間
	ShippingCostsWithinJapan *int       `form:"shippingCostsWithinJapan" json:"shippingCostsWithinJapan,omitempty"` // 初估日本國內運費
	GrossWeight              *int       `form:"grossWeight" json:"grossWeight"`                                     // 實際重量
	VolumetricWeight         *int       `form:"volumetricWeight" json:"volumetricWeight"`                           // 體積重量
	Status                   *uint8     `form:"status" json:"status"`
	CreatedAt                *time.Time `json:"createdAt"`
	UpdatedAt                *time.Time `json:"updatedAt"`
}

type ItemPhoto struct {
	ItemId    *uint64    `form:"itemId" validate:"required,gt=0" json:"-"`
	Sorted    *uint8     `form:"sorted" validate:"required,gt=0" json:"sorted"`
	Photo     *string    `form:"photo" validate:"required" json:"photo"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type ItemPastStatus struct {
	ItemId    *uint64    `form:"itemId" validate:"required,gt=0" json:"-"`
	Status    *uint8     `form:"status" validate:"required,gt=0" json:"status"`
	CreatedAt *time.Time `form:"createdAt" validate:"required" json:"createdAt"`
}

type ReorderItemPhotoReq struct {
	OriginalSorted uint8 `form:"originalSorted" validate:"required"`
	NewSorted      uint8 `form:"newSorted" validate:"required"`
}

type GetItemsReq struct {
	ConsignorId uint64    `form:"consignorId"` // 如果改成多個，做快取會有刪不到的風險
	WarehouseId string    `form:"warehouseId"`
	ExpireAt    time.Time `form:"expireAt"`
	Status      []uint8   `form:"status"`
	StartAt     time.Time `form:"startAt"`
	EndAt       time.Time `form:"endAt"`
	Sort        []string  `form:"sort"`
	Order       []string  `form:"order"`
	Limit       int64     `form:"limit"`
	Offset      int64     `form:"offset"`
}

type GetItemsRes struct {
	Items []Item `json:"items"`
	Count int64  `json:"count"`
}
