package workers

import (
	"fmt"
	"time"
)

const (
	// 出價後狀態
	HighestBiddingStatus   uint8 = 1 // 系統出價最高者
	NotHighestBiddedStatus uint8 = 2 // 系統已出價但未最高者
)

var (
	ErrLoginFailed = fmt.Errorf("Login failed")
)

type Cookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Path     string `json:"path,omitempty"`
	Domain   string `json:"domain,omitempty"`
	Expires  int64  `json:"expirationDate,omitempty"`
	Secure   bool   `json:"secure,omitempty"`
	HttpOnly bool   `json:"httpOnly,omitempty"`
	SameSite string `json:"sameSite,omitempty"`
}

type GetAuctionItemInfo struct {
	AuctionItemInfo
	Bidders   []Bidder `json:"bidders"`
	Countdown int      `json:"countdown"`
}

type AuctionItemInfo struct {
	IsClosed     bool      `json:"isClosed"`
	SellerName   string    `json:"sellerName"`
	Name         string    `json:"name"`
	Photo        string    `json:"photo"`
	CurrentPrice int       `json:"currentPrice"`
	CloseAt      time.Time `json:"closeAt"`
}

type Bidder struct {
	Account   string    `json:"account"`
	Rating    int       `json:"rating"`
	BidAmount int       `json:"bidAmount"`
	Quantity  int       `json:"quantity"`
	LastBidAt time.Time `json:"lastBidAt"`
}
