package consignors

import "time"

var (
	// Consignor 狀態
	EnabledStatus uint8 = 1 // 啟用中

	AwaitingVerificationCompletionStatus uint8 = 11 // 身份驗證中

	DisabledStatus uint8 = 99 // 禁用中
)

type Consignor struct {
	ID             *uint64    `json:"id"`
	Account        *string    `form:"account" json:"account"`
	Password       *string    `form:"password" json:"password"`
	Nickname       *string    `form:"nickname" json:"nickname"`
	Name           *string    `form:"name" json:"name"`
	Identification *string    `form:"identification" json:"identification"`
	BankCode       *string    `form:"bankCode" json:"bankCode"`
	BankAccount    *string    `form:"bankAccount" json:"bankAccount"`
	CommissionRate *float64   `form:"commissionRate" json:"commissionRate"`
	Status         *uint8     `form:"status" json:"status"`
	CreatedAt      *time.Time `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
}