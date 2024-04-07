package errno

// 1000 ~7999 為各服務錯誤碼，如果多個服務都會用到的錯誤碼，那就定義在 8000

const (
	// -------------------- 10xx ~ 89xx 內部服務 --------------------

	// 1000 auction-master
	AuctionMaster        = "1000"
	WorkerIdDoesNotExist = "1001"

	// 1100 jobs
	Jobs = "1100"

	// 1200 auction-items
	AuctionItems = "1200"

	// 1300 workers
	Workers = "1300"

	// 1400 worker
	Worker = "1400"

	// -------------------- 90xx ~ 99xx 外部服務 --------------------

)