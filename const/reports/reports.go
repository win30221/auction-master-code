package reports

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	SoldItemType = 1000 // 賣出物品

	PayYahooFeeType                   = 2000 // 支付結標日拍手續費
	PayAuctionItemCancellationFeeType = 2001 // 支付取消日拍手續費
	PaySpaceFeeType                   = 2002 // 支付留倉費
	PayShippingCostType               = 2003 // 支付運費
)

type Report struct {
	ID                *primitive.ObjectID `json:"id" bson:"_id"`
	Type              *int                `form:"type" json:"type" bson:"type" sendForm:"type"`
	ConsignorID       *int                `form:"consignorID" json:"consignorID" bson:"consignorID" sendForm:"consignorID"`
	ConsignorNickname *string             `form:"consignorNickname" json:"consignorNickname" bson:"consignorNickname" sendForm:"consignorNickname"`
	OpCode            *string             `form:"opCode" json:"opCode" bson:"opCode" sendForm:"opCode"`
	ItemID            *int                `form:"itemID" json:"itemID" bson:"itemID" sendForm:"itemID"`
	AuctionID         *int                `form:"auctionID" json:"auctionID,omitempty" bson:"auctionID,omitempty" sendForm:"auctionID"`
	Price             *int                `form:"price" json:"price,omitempty" bson:"price,omitempty" sendForm:"price"`
	Currency          *string             `form:"currency" json:"currency,omitempty" bson:"currency,omitempty" sendForm:"currency"`
	YahooFee          *int                `form:"yahooFee" json:"yahooFee,omitempty" bson:"yahooFee,omitempty" sendForm:"yahooFee"`
	Commission        *int                `form:"commission" json:"commission,omitempty" bson:"commission,omitempty" sendForm:"commission"`
	Bonus             *int                `form:"bonus" json:"bonus,omitempty" bson:"bonus,omitempty" sendForm:"bonus"`
	SpaceFee          *int                `form:"spaceFee" json:"spaceFee,omitempty" bson:"spaceFee,omitempty" sendForm:"spaceFee"`
	ShippingCost      *int                `form:"shippingCost" json:"shippingCost,omitempty" bson:"shippingCost,omitempty" sendForm:"shippingCost"`
	CreatedAt         *time.Time          `json:"createdAt" bson:"createdAt"`
}

type GetReportsRes struct {
	Reports []Report `json:"reports"`
	Count   int64    `json:"count"`
}
