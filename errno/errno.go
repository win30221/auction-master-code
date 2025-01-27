package errno

// 1000 ~7999 為各服務錯誤碼，如果多個服務都會用到的錯誤碼，那就定義在 8000

const (
	// -------------------- 10xx ~ 89xx 內部服務 --------------------

	// 1000 auction-master
	AuctionMaster           = "1000"
	PermissionDenied        = "1001" // 無權限
	BlockedUser             = "1002" // 登入錯誤
	TokenIncorrect          = "1003" // Token 錯誤
	PasswordIncorrect       = "1004" // 密碼錯誤
	NameIncorrect           = "1005" // 名字錯誤
	IdentificationIncorrect = "1006" // 身份證錯誤

	WorkerIdDoesNotExist         = "1100" // 查無 worker id
	ItemOnBidding                = "1101" // 物品競標中
	ItemExpired                  = "1102" // 物品過期
	ItemTypeNotSet               = "1103" // 物品類型未設定
	MinEstimatedPriceNotSet      = "1104" // 最低估價未設定
	AuctionItemNotClosed         = "1105" // 日拍競標商品未結標
	ConsignorVerified            = "1106" // 使用者已驗證
	ConsignorUnverified          = "1107" // 使用者未驗證
	ConsignorAccountExists       = "1108" // 使用者帳號已存在
	ConsignorNicknameExists      = "1109" // 使用者暱稱已存在
	ExceedCancellationBufferTime = "1110" // 超過取消緩衝時間
	SameWarehouseId              = "1111" // 相同倉庫id
	InvalidAccountFormat         = "1112" // 帳號規則錯誤
	NoItemIdsProvided            = "1113" // 請提供 item id
	InvalidDateRange             = "1114" // 超出時間範圍
	QueryLimitExceeded           = "1115" // 查詢已達上限
	PhotoLimitExceeded           = "1116" // 照片已達上限
	LineIdBindingConflict        = "1117" // Line 綁定衝突
	LineIdBindingLinkExpired     = "1118" // Line 綁定連結已過期

	// 3000 jobs
	Jobs = "3000"

	// 3100 admins
	Admins        = "3100"
	AdminExisted  = "3101" // 使用者已存在
	AdminNotExist = "3102" // 使用者不存在

	// 3200 consignors
	Consignors                    = "3200"
	ConsignorExisted              = "3201" // 使用者已存在
	ConsignorNotExist             = "3202" // 使用者不存在
	ConsignorVerificationExisted  = "3203" // 審核申請已存在
	ConsignorVerificationNotExist = "3204" // 審核申請不存在

	// 3300 items
	Items        = "3300"
	ItemNotExist = "3301" // 物品不存在

	// 3400 wallets
	Wallets                   = "3400"
	WalletExisted             = "3401" // 錢包已存在
	WalletNotExist            = "3402" // 錢包不存在
	WalletInsufficientBalance = "3403" // 錢包餘額不足

	// 3500 bonuses
	Bonuses                  = "3500"
	BonusExisted             = "3501" // 紅利已存在
	BonusNotExist            = "3502" // 紅利不存在
	BonusInsufficientBalance = "3503" // 紅利餘額不足

	// 3600 auction-items
	AuctionItems        = "3600" // 日拍競標商品錯誤
	AuctionItemExist    = "3601" // 日拍競標商品已存在
	AuctionItemNotExist = "3602" // 日拍競標商品不存在

	// 3700 workers
	Workers = "3700"

	// 3800 worker
	Worker            = "3800"
	WorkerNotLoggedIn = "3801" // worker 未登入

	// 3900 shippings
	Shippings        = "3900"
	ShippingNotExist = "3901" // 出貨單不存在

	// 4000 reports
	Reports        = "4000"
	ReportNotExist = "4001" // 報表不存在
	RecordNotExist = "4002" // 交易紀錄不存在

	// -------------------- 90xx ~ 99xx 外部服務 --------------------

)
