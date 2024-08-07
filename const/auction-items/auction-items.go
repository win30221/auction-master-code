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
	AwaitingConsignorPayFeeStatus      uint8 = 12 // 等待寄售人付款
	ConsignorRequestCancellationStatus uint8 = 13 // 寄售人申請取消

	SoldStatus             uint8 = 21 // 售出
	CanceledStatus         uint8 = 22 // 手動取消
	ConsignorFeePaidStatus uint8 = 23 // 寄售人已付手續費

)

// 競標商品 schema
type AuctionItem struct {
	ID           *uint64    `json:"id"`
	ConsignorID  *uint64    `form:"consignorID" json:"consignorID" sendForm:"consignorID"`
	ItemID       *uint64    `form:"itemID" json:"itemID" sendForm:"itemID"`
	SellerID     *uint64    `form:"sellerID" json:"sellerID" sendForm:"sellerID"`
	WatcherID    *uint64    `form:"watcherID" json:"watcherID" sendForm:"watcherID"`
	AuctionID    *string    `form:"auctionID" json:"auctionID" sendForm:"auctionID"`
	Name         *string    `form:"name" json:"name" sendForm:"name"`
	Photo        *string    `form:"photo" json:"photo" sendForm:"photo"`
	ReservePrice *int       `form:"reservePrice" json:"reservePrice" sendForm:"reservePrice"`
	CurrentPrice *int       `form:"currentPrice" json:"currentPrice" sendForm:"currentPrice"`
	HighestPrice *int       `form:"highestPrice" json:"highestPrice" sendForm:"highestPrice"`
	CloseAt      *time.Time `form:"closeAt" json:"closeAt" sendForm:"closeAt"`
	ClosedPrice  *int       `form:"closedPrice" json:"closedPrice" sendForm:"closedPrice"`
	Status       *uint8     `form:"status" json:"status" sendForm:"status"`
	CreatedAt    *time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt    *time.Time `form:"updatedAt" json:"updatedAt"`
}

type GetAuctionItemsRes struct {
	AuctionItems []AuctionItem `json:"auctionItems"`
	Count        int           `json:"count"`
}
