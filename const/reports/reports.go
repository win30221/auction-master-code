package reports

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	SoldItemType              int = 1000 // 賣出物品
	CompanyDirectPurchaseType int = 1001 // 公司直購
	CompanyPurchasedType      int = 1002 // 公司買回

	PayYahooFeeType                   int = 2000 // 支付結標日拍手續費
	PayAuctionItemCancellationFeeType int = 2001 // 支付取消日拍手續費
	PaySpaceFeeType                   int = 2002 // 支付留倉費
	PayShippingCostType               int = 2003 // 支付運費

	UnpaidStatus          uint8 = 1  // 未付款
	SubmitAppraisalStatus uint8 = 2  // 已提交付款
	PaidStatus            uint8 = 10 // 已付款
)

type Summary struct {
	TotalClosedPrice          int `json:"totalClosedPrice" bson:"totalClosedPrice"`
	TotalPrice                int `json:"totalPrice" bson:"totalPrice"`
	TotalDirectPurchasePrice  int `json:"totalDirectPurchasePrice" bson:"totalDirectPurchasePrice"`
	TotalPurchasedPrice       int `json:"totalPurchasedPrice" bson:"totalPurchasedPrice"`
	TotalYahooFee             int `json:"totalYahooFee" bson:"totalYahooFee"`
	TotalCommission           int `json:"totalCommission" bson:"totalCommission"`
	TotalBonus                int `json:"totalBonus" bson:"totalBonus"`
	TotalProfit               int `json:"totalProfit" bson:"totalProfit"`
	TotalYahooCancellationFee int `json:"totalYahooCancellationFee"`
	TotalSpaceFee             int `json:"totalSpaceFee" bson:"totalSpaceFee"`
	TotalShippingCost         int `json:"totalShippingCost" bson:"totalShippingCost"`
}

type Report struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id"`
	Reports   *map[string]Summary `json:"reports" bson:"reports"`
	ReportAt  *time.Time          `json:"reportAt" bson:"reportAt"`
	CreatedAt *time.Time          `json:"createdAt" bson:"createdAt"`
}

type Record struct {
	ID                   *primitive.ObjectID `json:"id" bson:"_id"`
	Type                 *int                `form:"type" json:"type" bson:"type" sendForm:"type"`
	ConsignorID          *uint64             `form:"consignorID" json:"consignorID" bson:"consignorID" sendForm:"consignorID"`
	ConsignorNickname    *string             `form:"consignorNickname" json:"consignorNickname" bson:"consignorNickname" sendForm:"consignorNickname"`
	OpCode               *string             `form:"opCode" json:"opCode" bson:"opCode" sendForm:"opCode"`
	ItemID               *uint64             `form:"itemID" json:"itemID" bson:"itemID" sendForm:"itemID"`
	AuctionID            *uint64             `form:"auctionID" json:"auctionID,omitempty" bson:"auctionID,omitempty" sendForm:"auctionID"`
	Currency             *string             `form:"currency" json:"currency,omitempty" bson:"currency,omitempty" sendForm:"currency"`
	ExchangeRate         *float64            `form:"exchangeRate" json:"exchangeRate,omitempty" bson:"exchangeRate,omitempty" sendForm:"exchangeRate"`
	ClosedPrice          *int                `form:"closedPrice" json:"closedPrice,omitempty" bson:"closedPrice,omitempty" sendForm:"closedPrice"`
	Price                *int                `form:"price" json:"price,omitempty" bson:"price,omitempty" sendForm:"price"`
	DirectPurchasePrice  *int                `form:"directPurchasePrice" json:"directPurchasePrice,omitempty" bson:"directPurchasePrice,omitempty" sendForm:"directPurchasePrice"`
	PurchasedPrice       *int                `form:"purchasedPrice" json:"purchasedPrice,omitempty" bson:"purchasedPrice,omitempty" sendForm:"purchasedPrice"`
	YahooFeeRate         *float64            `form:"yahooFeeRate" json:"yahooFeeRate,omitempty" bson:"yahooFeeRate,omitempty" sendForm:"yahooFeeRate"`
	YahooFee             *int                `form:"yahooFee" json:"yahooFee,omitempty" bson:"yahooFee,omitempty" sendForm:"yahooFee"`
	CommissionRate       *float64            `form:"commissionRate" json:"commissionRate,omitempty" bson:"commissionRate,omitempty" sendForm:"commissionRate"`
	Commission           *int                `form:"commission" json:"commission,omitempty" bson:"commission,omitempty" sendForm:"commission"`
	BonusRate            *float64            `form:"bonusRate" json:"bonusRate,omitempty" bson:"bonusRate,omitempty" sendForm:"bonusRate"`
	Bonus                *int                `form:"bonus" json:"bonus,omitempty" bson:"bonus,omitempty" sendForm:"bonus"`
	Profit               *int                `form:"profit" json:"profit,omitempty" bson:"profit,omitempty" sendForm:"profit"`
	YahooCancellationFee *int                `form:"yahooCancellationFee" json:"yahooCancellationFee,omitempty" bson:"yahooCancellationFee,omitempty" sendForm:"yahooCancellationFee"`
	SpaceFee             *int                `form:"spaceFee" json:"spaceFee,omitempty" bson:"spaceFee,omitempty" sendForm:"spaceFee"`
	ShippingCost         *int                `form:"shippingCost" json:"shippingCost,omitempty" bson:"shippingCost,omitempty" sendForm:"shippingCost"`
	Status               *uint8              `form:"status" json:"status" sendForm:"status"`
	CreatedAt            *time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt            *time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type CreateReportReq struct {
	ReportAt time.Time `form:"reportAt" validate:"required"`
}

type GetReportsReq struct {
	StartAt time.Time `form:"startAt"`
	EndAt   time.Time `form:"endAt"`
}

type GetRecordsReq struct {
	Type        []int     `form:"type"`
	ConsignorID []uint64  `form:"consignorID"`
	Status      []int     `form:"status"` // mongo 查詢沒有 uint8 類型
	StartAt     time.Time `form:"startAt"`
	EndAt       time.Time `form:"endAt"`
	Limit       int       `form:"limit"`
	Offset      int       `form:"offset"`
}

type GetRecordsRes struct {
	Records []Record `json:"records"`
	Count   int64    `json:"count"`
}
