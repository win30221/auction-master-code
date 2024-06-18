package opcode

const (
	Deal                    = "Deal-Item:%d-AuctionItem:%s"
	PayFee                  = "PayFee-Item:%d-AuctionItem:%s"
	PaySpaceFee             = "PaySpaceFee-Item:%d"
	PayCancelAuctionItemFee = "PayCancelAuctionItemFee-AuctionItem:%s"
)

const (
	WalletSoldType = "1000"

	WalletPayFeeType              = "2000"
	WalletPaySpaceFeeType         = "2001"
	WalletPayCancelAuctionItemFee = "2002"
)

const (
	BonusSoldType = "1000"
)
