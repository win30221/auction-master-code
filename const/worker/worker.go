package worker

import (
	"time"
)

var (
	// 競標商品狀態
	YahooAuctionHighestBidded    uint8 = 1 // 系統出價最高者
	YahooAuctionNotHighestBidded uint8 = 2 // 系統已出價但未最高者
)

type Cookie struct {
	Name     string  `json:"name"`
	Value    string  `json:"value"`
	Path     string  `json:"path,omitempty"`
	Domain   string  `json:"domain,omitempty"`
	Expires  float64 `json:"expirationDate,omitempty"`
	Secure   bool    `json:"secure,omitempty"`
	HttpOnly bool    `json:"httpOnly,omitempty"`
	SameSite string  `json:"sameSite,omitempty"`
}

type AuctionItemInfo struct {
	AuctionItemBasicInfo
	Bidders    []Bidder  `json:"bidders"`
	Countdown  int       `json:"countdown"`
	RetrieveAt time.Time `json:"retrieveAt"`
}

type AuctionItemBasicInfo struct {
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
