package auctionitems

import (
	"time"
)

const (
	// 競標商品狀態
	WatchingStatus uint8 = 1 // 盯標中

	HighestBiddedStatus    uint8 = 11 // 系統出價最高者
	NotHighestBiddedStatus uint8 = 12 // 系統已出價但未最高者
	StopWatchingStatus     uint8 = 13 // 停止盯標
	ClosedStatus           uint8 = 14 // 結標

	SoldStatus               uint8 = 21 // 售出
	CompanyRepurchasedStatus uint8 = 22 // 被公司買回
	CanceledStatus           uint8 = 23 // 手動取消
)

// 競標商品 schema
type AuctionItem struct {
	ID           int64     `form:"id" json:"id"`
	ItemID       int64     `form:"itemID" json:"itemID"`
	SellerID     string    `form:"sellerID" json:"sellerID"`
	WatcherID    string    `form:"watcherID" json:"watcherID"`
	AuctionID    string    `form:"auctionID" json:"auctionID"`
	Name         string    `form:"name" json:"name"`
	Photo        string    `form:"photo" json:"photo"`
	ReservePrice int       `form:"reservePrice" json:"reservePrice"`
	CurrentPrice int       `form:"currentPrice" json:"currentPrice"`
	HighestPrice int       `form:"highestPrice" json:"highestPrice"`
	CloseAt      time.Time `form:"closeAt" json:"closeAt"`
	ClosedPrice  int       `form:"closedPrice" json:"closedPrice"`
	Status       uint8     `form:"status" json:"status"`
	CreatedAt    time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time `form:"updatedAt" json:"updatedAt"`
}
