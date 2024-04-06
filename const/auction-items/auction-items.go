package auctionitems

import (
	"time"
)

const (
	// 競標商品狀態
	InitStatus        uint8 = 1 // 初始化
	StopBiddingStatus uint8 = 2 // 停止出價

	ClosedStatus uint8 = 11 // 結標

	SoldStatus               uint8 = 21 // 售出
	CanceledStatus           uint8 = 22 // 手動取消
	CompanyRepurchasedStatus uint8 = 23 // 被公司買回
)

// 競標商品 schema
type AuctionItem struct {
	ID           *int64     `form:"id" json:"id"`
	ItemID       *int64     `form:"itemID" json:"itemID"`
	SellerID     *int64     `form:"sellerID" json:"sellerID"`
	WatcherID    *int64     `form:"watcherID" json:"watcherID"`
	AuctionID    *string    `form:"auctionID" json:"auctionID"`
	Name         *string    `form:"name" json:"name"`
	Photo        *string    `form:"photo" json:"photo"`
	ReservePrice *int       `form:"reservePrice" json:"reservePrice"`
	CurrentPrice *int       `form:"currentPrice" json:"currentPrice"`
	HighestPrice *int       `form:"highestPrice" json:"highestPrice"`
	CloseAt      *time.Time `form:"closeAt" json:"closeAt"`
	ClosedPrice  *int       `form:"closedPrice" json:"closedPrice"`
	Status       *uint8     `form:"status" json:"status"`
	CreatedAt    *time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt    *time.Time `form:"updatedAt" json:"updatedAt"`
}
