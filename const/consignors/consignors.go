package consignors

import "time"

var (
	// Consignor 狀態
	EnabledStatus uint8 = 1 // 啟用中

	AwaitingVerificationCompletionStatus uint8 = 11 // 身份尚未驗證
	VerificationSuccessfulStatus         uint8 = 12 // 驗證成功
	VerificationFailedStatus             uint8 = 13 // 驗證失敗

	DisabledStatus uint8 = 99 // 禁用中
)

type Consignor struct {
	ID             *uint64    `json:"id"`
	Account        *string    `form:"account" json:"account"`
	Password       *string    `form:"password" json:"password"`
	Nickname       *string    `form:"nickname" json:"nickname"`
	Name           *string    `form:"name" json:"name"`
	Identification *string    `form:"identification" json:"identification"`
	Phone          *string    `form:"phone" json:"phone"`
	BankCode       *string    `form:"bankCode" json:"bankCode"`
	BankAccount    *string    `form:"bankAccount" json:"bankAccount"`
	Status         *uint8     `form:"status" json:"status"`
	CreatedAt      *time.Time `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
}

type ConsignorVerification struct {
	ID             *uint64    `json:"id"`
	ConsignorID    *uint64    `form:"consignorID" json:"consignorID"`
	Photo          *string    `form:"photo" json:"photo"`
	Name           *string    `form:"name" json:"name"`
	Identification *string    `form:"identification" json:"identification"`
	Phone          *string    `form:"phone" json:"phone"`
	BankCode       *string    `form:"bankCode" json:"bankCode"`
	BankAccount    *string    `form:"bankAccount" json:"bankAccount"`
	Status         *uint8     `form:"status" json:"status"`
	CreatedAt      *time.Time `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
}

type GetConsignorsRes struct {
	Consignors []Consignor `json:"consignors"`
	Count      int         `json:"count"`
}

type GetConsignorVerificationsRes struct {
	ConsignorVerifications []ConsignorVerification `json:"consignorVerifications"`
	Count                  int                     `json:"count"`
}
