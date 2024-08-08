package opcode

const (
	Deal                          = "Deal-Item:%d-AuctionItem:%s"
	PayFee                        = "PayFee-Item:%d-AuctionItem:%s"
	PaySpaceFee                   = "PaySpaceFee-Item:%s" // PaySpaceFee-Item:1-2
	PayAuctionItemCancellationFee = "PayAuctionItemCancellationFee-AuctionItem:%s"
)

const (
	WalletSoldType = "Sold"

	WalletPayFeeType                        = "PayFee"
	WalletPaySpaceFeeType                   = "PaySpaceFee"
	WalletPayAuctionItemCancellationFeeType = "PayAuctionItemCancellationFee"
)

const (
	BonusSoldType = "Sold"
)
