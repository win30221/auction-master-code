package auctionitems

import (
	"time"
)

var (
	// 競標商品狀態
	InitStatus             uint8 = 1 // 初始化
	StopBiddingStatus      uint8 = 2 // 停止出價
	HighestBiddedStatus    uint8 = 3 // 系統出價最高者
	NotHighestBiddedStatus uint8 = 4 // 系統已出價但未最高者

	ClosedStatus                       uint8 = 11 // 結標
	ConsignorRequestCancellationStatus uint8 = 12 // 寄售人申請取消
	AwaitingConsignorPayFeeStatus      uint8 = 13 // 等待寄售人付款

	SoldStatus             uint8 = 21 // 售出
	CanceledStatus         uint8 = 22 // 手動取消
	ConsignorFeePaidStatus uint8 = 23 // 寄售人已付手續費
	NoBidsPlacedStatus     uint8 = 24 // 無人下標
)

// 競標商品 schema
type AuctionItem struct {
	AuctionId    *string    `form:"auctionId" json:"auctionId"`
	ConsignorId  *uint64    `form:"consignorId" json:"consignorId"`
	ItemId       *uint64    `form:"itemId" json:"itemId"`
	SellerId     *uint64    `form:"sellerId" json:"sellerId"`
	WatcherId    *uint64    `form:"watcherId" json:"watcherId"`
	Name         *string    `form:"name" json:"name"`
	Photo        *string    `form:"photo" json:"photo"`
	ReservePrice *int       `form:"reservePrice" json:"reservePrice"`
	CurrentPrice *int       `form:"currentPrice" json:"currentPrice"`
	HighestPrice *int       `form:"highestPrice" json:"highestPrice"`
	CloseAt      *time.Time `form:"closeAt" json:"closeAt"`
	ClosedPrice  *int       `form:"closedPrice" json:"closedPrice"`
	Status       *uint8     `form:"status" json:"status"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
}

type GetAuctionItemsReq struct {
	ConsignorId uint64    `form:"consignorId"` // 如果改成多個，做快取會有刪不到的風險
	ItemId      []uint64  `form:"itemId"`
	Status      []uint8   `form:"status"`
	StartAt     time.Time `form:"startAt"`
	EndAt       time.Time `form:"endAt"`
	Sort        []string  `form:"sort"`
	Order       []string  `form:"order"`
	Limit       int64     `form:"limit"`
	Offset      int64     `form:"offset"`
}

type GetAuctionItemsRes struct {
	AuctionItems []AuctionItem `json:"auctionItems"`
	Count        int64         `json:"count"`
}
