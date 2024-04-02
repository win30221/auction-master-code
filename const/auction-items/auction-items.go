package auctionitems

import (
	"time"

	"github.com/win30221/auction-master-code/const/workers"
)

const (

	// 競標商品狀態
	Watching         uint8 = 1  // 盯標中
	HighestBidded    uint8 = 11 // 系統出價最高者
	NotHighestBidded uint8 = 12 // 系統已出價但未最高者
	StopWatching     uint8 = 13 // 停止盯標
	Closed           uint8 = 21 // 商品關閉時已達預期金額
	Canceled         uint8 = 22 // 商品關閉時未達預期金額
)

type AuctionItem struct {
	ID           int64     `form:"id" json:"id"`
	ItemID       int64     `form:"itemID" json:"itemID"`
	SellerID     string    `form:"sellerID" json:"sellerID"`
	WatcherID    string    `form:"watcherID" json:"watcherID"`
	AuctionID    string    `form:"auctionID" json:"auctionID"`
	AuctionName  string    `form:"auctionName" json:"auctionName"` // 有 itemID 後就移除
	AuctionPhoto string    `form:"auctionPhoto" json:"auctionPhoto"`
	TargetPrice  int       `form:"targetPrice" json:"targetPrice"` // 有 itemID 後就移除
	CurrentPrice int       `form:"currentPrice" json:"currentPrice"`
	HighestPrice int       `form:"highestPrice" json:"highestPrice"`
	CloseAt      time.Time `form:"closeAt" json:"closeAt"`
	ClosedPrice  int       `form:"closedPrice" json:"closedPrice"`
	Status       uint8     `form:"status" json:"status"`
	CreatedAt    time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time `form:"updatedAt" json:"updatedAt"`
}

type WatchedAuctionItem struct {
	ID           int64                `json:"id"`
	SellerName   string               `json:"sellerName"`
	WatcherName  string               `json:"watcherName"`
	OwnerName    string               `json:"ownerName"`
	AuctionID    string               `json:"auctionID"`
	AuctionName  string               `json:"auctionName"`
	AuctionPhoto string               `json:"auctionPhoto"`
	Bidders      []workers.BidderInfo `json:"bidders"`
	TargetPrice  int                  `json:"targetPrice"`
	HighestPrice int                  `json:"highestPrice"`
	CloseAt      time.Time            `json:"closeAt"`
	ClosedPrice  int                  `json:"closedPrice"`
	Status       uint8                `json:"status"`
	CreatedAt    time.Time            `json:"createdAt"`
	UpdatedAt    time.Time            `json:"updatedAt"`
}
