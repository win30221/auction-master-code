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
	Avatar              *string    `form:"avatar" json:"avatar"`
	Account             *string    `form:"account" json:"account"`
	Password            *string    `form:"password" json:"password"`
	Nickname            *string    `form:"nickname" json:"nickname"`
	CommissionBonusRate *float64   `form:"commissionBonusRate" json:"commissionBonusRate"`
	Name                *string    `form:"name" json:"name"`
	Identification      *string    `form:"identification" json:"identification"`
	Gender              *uint8     `form:"gender" json:"gender"`
	Birthday            *time.Time `form:"birthday" json:"birthday"`
	City                *string    `form:"city" json:"city"`
	District            *string    `form:"district" json:"district"`
	StreetAddress       *string    `form:"streetAddress" json:"streetAddress"`
	Phone               *string    `form:"phone" json:"phone"`
	BankCode            *string    `form:"bankCode" json:"bankCode"`
	BankAccount         *string    `form:"bankAccount" json:"bankAccount"`
	Status              *uint8     `form:"status" json:"status"`
	CreatedAt           *time.Time `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt"`
}

type ConsignorVerification struct {
	ID             *uint64    `json:"id"`
	ConsignorID    *uint64    `form:"consignorID" json:"consignorID"`
	Photo          *string    `form:"photo" json:"photo"`
	Name           *string    `form:"name" json:"name"`
	Identification *string    `form:"identification" json:"identification"`
	Gender         *uint8     `form:"gender" json:"gender"`
	Birthday       *time.Time `form:"birthday" json:"birthday"`
	City           *string    `form:"city" json:"city"`
	District       *string    `form:"district" json:"district"`
	StreetAddress  *string    `form:"streetAddress" json:"streetAddress"`
	Phone          *string    `form:"phone" json:"phone"`
	BankCode       *string    `form:"bankCode" json:"bankCode"`
	BankAccount    *string    `form:"bankAccount" json:"bankAccount"`
	Status         *uint8     `form:"status" json:"status"`
	CreatedAt      *time.Time `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
}

type GetConsignorsReq struct {
	FuzzyNickname string   `form:"fuzzyNickname"`
	Status        []uint8  `form:"status"`
	Sort          []string `form:"sort"`
	Order         []string `form:"order"`
	Limit         int64    `form:"limit"`
	Offset        int64    `form:"offset"`
}

type GetConsignorsRes struct {
	Consignors []Consignor `json:"consignors"`
	Count      int         `json:"count"`
}

type GetConsignorVerificationsReq struct {
	ConsignorID uint64   `form:"consignorID"` // 如果改成多個，做快取會有刪不到的風險
	Status      []uint8  `form:"status"`
	Sort        []string `form:"sort"`
	Order       []string `form:"order"`
	Limit       int64    `form:"limit"`
	Offset      int64    `form:"offset"`
}

type GetConsignorVerificationsRes struct {
	ConsignorVerifications []ConsignorVerification `json:"consignorVerifications"`
	Count                  int                     `json:"count"`
}
