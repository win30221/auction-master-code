package workers

import (
	"fmt"
	"time"
)

const (
	// 出價後狀態
	HighestBiddingStatus   int = 1 // 系統出價最高者
	NotHighestBiddedStatus int = 2 // 系統已出價但未最高者
	HasBeenBlocked         int = 3 // 被賣家封鎖
)

var (
	ErrLoginFailed = fmt.Errorf("Login failed")
)

type BidderInfo struct {
	Account   string    `json:"account"`
	Rating    int       `json:"rating"`
	BidAmount int       `json:"bidAmount"`
	Quantity  int       `json:"quantity"`
	LastBidAt time.Time `json:"lastBidAt"`
}

type AuctionItemInfo struct {
	IsClosed     bool      `json:"isClosed"`
	Name         string    `json:"name"`
	Photo        string    `json:"photo"`
	CurrentPrice int       `json:"currentPrice"`
	CloseAt      time.Time `json:"closeAt"`
}

type GetAuctionItemInfo struct {
	AuctionItemInfo
	BidderInfo []BidderInfo `json:"bidderInfo"`
	Countdown  int          `json:"countdown"`
}
