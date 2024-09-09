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
	ID                       *uint64    `form:"id" json:"id"`
	ConsignorID              *uint64    `form:"consignorID" json:"consignorID" sendForm:"consignorID"`
	Type                     *uint8     `form:"type" json:"type" sendForm:"type"`                                                                       // 物品類型
	IsNew                    *bool      `form:"isNew" json:"isNew"  sendForm:"isNew"`                                                                   // 是否全新品
	Name                     *string    `form:"name" json:"name" sendForm:"name"`                                                                       // 物品名稱
	Description              *string    `form:"description" json:"description" sendForm:"description"`                                                  // 物品描述
	DirectPurchasePrice      *int       `form:"directPurchasePrice" json:"directPurchasePrice" sendForm:"directPurchasePrice"`                          // 直購金額
	MinEstimatedPrice        *int       `form:"minEstimatedPrice" json:"minEstimatedPrice" sendForm:"minEstimatedPrice"`                                // 最低估價金額
	MaxEstimatedPrice        *int       `form:"maxEstimatedPrice" json:"maxEstimatedPrice" sendForm:"maxEstimatedPrice"`                                // 最高估價金額
	ReservePrice             *int       `form:"reservePrice" json:"reservePrice" sendForm:"reservePrice"`                                               // 期望金額(用於最後盯標)
	ExpireAt                 *time.Time `form:"expireAt" json:"expireAt" sendForm:"expireAt"`                                                           // 過期時間
	WarehouseID              *string    `form:"warehouseID" json:"warehouseID" sendForm:"warehouseID"`                                                  // 存放於倉庫的編號
	Space                    *uint8     `form:"space" json:"space" sendForm:"space"`                                                                    // 物品占用空間
	ShippingCostsWithinJapan *int       `form:"shippingCostsWithinJapan" json:"shippingCostsWithinJapan,omitempty" sendForm:"shippingCostsWithinJapan"` // 初估日本國內運費
	GrossWeight              *int       `form:"grossWeight" json:"grossWeight" sendForm:"grossWeight"`                                                  // 實際重量
	VolumetricWeight         *int       `form:"volumetricWeight" json:"volumetricWeight" sendForm:"volumetricWeight"`                                   // 體積重量
	Status                   *uint8     `form:"status" json:"status" sendForm:"status"`
	CreatedAt                *time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt                *time.Time `form:"updatedAt" json:"updatedAt"`
}

type ItemPhoto struct {
	ItemID    *uint64    `form:"itemID" validate:"required,gt=0" json:"-" sendForm:"itemID"`
	Sorted    *uint8     `form:"sorted" validate:"required,gt=0" json:"sorted" sendForm:"sorted"`
	Photo     *string    `form:"photo" validate:"required" json:"photo" sendForm:"photo"`
	CreatedAt *time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time `form:"updatedAt" json:"updatedAt"`
}

type ItemPastStatus struct {
	ItemID    *uint64    `form:"itemID" validate:"required,gt=0" json:"-" sendForm:"itemID"`
	Status    *uint8     `form:"status" validate:"required,gt=0" json:"status" sendForm:"status"`
	CreatedAt *time.Time `form:"createdAt" validate:"required" json:"createdAt" sendForm:"createdAt"`
}

type ReorderItemPhotoReq struct {
	OriginalSorted uint8 `form:"originalSorted" validate:"required" sendForm:"originalSorted"`
	NewSorted      uint8 `form:"newSorted" validate:"required" sendForm:"newSorted"`
}

type GetItemsReq struct {
	ConsignorID uint64    `form:"consignorID" sendForm:"consignorID"` // 如果改成多個，做快取會有刪不到的風險
	WarehouseID string    `form:"warehouseID" sendForm:"warehouseID"`
	ExpireAt    time.Time `form:"expireAt" sendForm:"expireAt"`
	Status      []uint8   `form:"status" sendForm:"status"`
	StartAt     time.Time `form:"startAt" sendForm:"startAt"`
	EndAt       time.Time `form:"endAt" sendForm:"endAt"`
	Sort        []string  `form:"sort" sendForm:"sort"`
	Order       []string  `form:"order" sendForm:"order"`
	Limit       int64     `form:"limit" sendForm:"limit"`
	Offset      int64     `form:"offset" sendForm:"offset"`
}

type GetItemsRes struct {
	Items []Item `json:"items"`
	Count int    `json:"count"`
}
