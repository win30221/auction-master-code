package errno

// 1000 ~7999 為各服務錯誤碼，如果多個服務都會用到的錯誤碼，那就定義在 8000

const (
	// -------------------- 10xx ~ 89xx 內部服務 --------------------

	// 1000 auction-master
	AuctionMaster           = "1000"
	PermissionDenied        = "1001"
	BlockedUser             = "1002"
	TokenIncorrect          = "1003"
	PasswordIncorrect       = "1004"
	NameIncorrect           = "1005"
	IdentificationIncorrect = "1006"

	WorkerIdDoesNotExist         = "1020"
	ItemOnBidding                = "1021"
	ItemExpired                  = "1022"
	ItemTypeNotSet               = "1023"
	MinEstimatedPriceNotSet      = "1024"
	AuctionItemNotClosed         = "1025"
	ConsignorVerified            = "1026"
	ConsignorUnverified          = "1027"
	ConsignorAccountExists       = "1028"
	ConsignorNicknameExists      = "1029"
	ExceedCancellationBufferTime = "1030"

	// 1100 jobs
	Jobs = "1100"

	// 1200 auction-items
	AuctionItems        = "1200"
	AuctionItemExist    = "1201"
	AuctionItemNotExist = "1202"

	// 1300 workers
	Workers = "1300"

	// 1400 worker
	Worker            = "1400"
	WorkerNotLoggedIn = "1401"

	// 1500 admins
	Admins        = "1500"
	AdminExisted  = "1501" // 使用者已存在
	AdminNotExist = "1502" // 使用者不存在

	// 1600 consignors
	Consignors                    = "1600"
	ConsignorExisted              = "1601" // 使用者已存在
	ConsignorNotExist             = "1602" // 使用者不存在
	ConsignorVerificationExisted  = "1603" // 審核申請已存在
	ConsignorVerificationNotExist = "1604" // 審核申請不存在

	// 1700 wallets
	Wallets                   = "1700"
	WalletExisted             = "1701" // 錢包已存在
	WalletNotExist            = "1702" // 錢包不存在
	WalletInsufficientBalance = "1703" // 錢包餘額不足

	// 1800 items
	Items        = "1800"
	ItemNotExist = "1801"

	// 1900 bonuses
	Bonuses                  = "1900"
	BonusExisted             = "1901" // 紅利已存在
	BonusNotExist            = "1902" // 紅利不存在
	BonusInsufficientBalance = "1903" // 紅利餘額不足

	// 2000 shippings
	Shippings        = "2000"
	ShippingNotExist = "2001"

	// -------------------- 90xx ~ 99xx 外部服務 --------------------

)
