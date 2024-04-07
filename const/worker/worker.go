package worker

import (
	"time"
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
