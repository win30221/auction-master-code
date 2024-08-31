package consignors

import "time"

var (
	// Consignor 狀態
	EnabledStatus uint8 = 1 // 啟用中

	AwaitingVerificationCompletionStatus uint8 = 11 // 身份尚未驗證
	VerificationSuccessfulStatus         uint8 = 12 // 驗證成功
	VerificationFailedStatus             uint8 = 13 // 驗證失敗

	DisabledStatus uint8 = 99 // 禁用中

	MaleGender   uint8 = 1
	FemaleGender uint8 = 2
)

type Consignor struct {
	ID                  *uint64    `json:"id"`
	Avatar              *string    `form:"avatar" json:"avatar" sendForm:"avatar"`
	Account             *string    `form:"account" json:"account" sendForm:"account"`
	Password            *string    `form:"password" json:"password" sendForm:"password"`
	Nickname            *string    `form:"nickname" json:"nickname" sendForm:"nickname"`
	CommissionBonusRate *float64   `form:"commissionBonusRate" json:"commissionBonusRate" sendForm:"commissionBonusRate"`
	Name                *string    `form:"name" json:"name" sendForm:"name"`
	Identification      *string    `form:"identification" json:"identification" sendForm:"identification"`
	Gender              *uint8     `form:"gender" json:"gender" sendForm:"gender"`
	Birthday            *time.Time `form:"birthday" json:"birthday" sendForm:"birthday"`
	City                *string    `form:"city" json:"city" sendForm:"city"`
	District            *string    `form:"district" json:"district" sendForm:"district"`
	StreetAddress       *string    `form:"streetAddress" json:"streetAddress" sendForm:"streetAddress"`
	Phone               *string    `form:"phone" json:"phone" sendForm:"phone"`
	BankCode            *string    `form:"bankCode" json:"bankCode" sendForm:"bankCode"`
	BankAccount         *string    `form:"bankAccount" json:"bankAccount" sendForm:"bankAccount"`
	Status              *uint8     `form:"status" json:"status" sendForm:"status"`
	CreatedAt           *time.Time `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt"`
}

type ConsignorVerification struct {
	ID             *uint64    `json:"id"`
	ConsignorID    *uint64    `form:"consignorID" json:"consignorID" sendForm:"consignorID"`
	Photo          *string    `form:"photo" json:"photo" sendForm:"photo"`
	Name           *string    `form:"name" json:"name" sendForm:"name"`
	Identification *string    `form:"identification" json:"identification" sendForm:"identification"`
	Gender         *uint8     `form:"gender" json:"gender" sendForm:"gender"`
	Birthday       *time.Time `form:"birthday" json:"birthday" sendForm:"birthday"`
	City           *string    `form:"city" json:"city" sendForm:"city"`
	District       *string    `form:"district" json:"district" sendForm:"district"`
	StreetAddress  *string    `form:"streetAddress" json:"streetAddress" sendForm:"streetAddress"`
	Phone          *string    `form:"phone" json:"phone" sendForm:"phone"`
	BankCode       *string    `form:"bankCode" json:"bankCode" sendForm:"bankCode"`
	BankAccount    *string    `form:"bankAccount" json:"bankAccount" sendForm:"bankAccount"`
	Status         *uint8     `form:"status" json:"status" sendForm:"status"`
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
