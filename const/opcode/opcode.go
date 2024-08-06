package opcode

const (
	Deal                          = "Deal-Item:%d-AuctionItem:%s"
	PayFee                        = "PayFee-Item:%d-AuctionItem:%s"
	PaySpaceFee                   = "PaySpaceFee-Item:%s" // PaySpaceFee-Item:1-2
	PayAuctionItemCancellationFee = "PayAuctionItemCancellationFee-AuctionItem:%s"
)

const (
	WalletSoldType = "1000"

	WalletPayFeeType                        = "2000"
	WalletPaySpaceFeeType                   = "2001"
	WalletPayAuctionItemCancellationFeeType = "2002"
)

const (
	BonusSoldType = "1000"
)
