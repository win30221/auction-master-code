package workers

import (
	"time"
)

var (
	// Worker 類型
	SellerType  = "Seller"
	WatcherType = "Watcher"

	// Worker 狀態
	ActiveStatus                  uint8 = 1  // 啟用中
	AwaitingSetupCompletionStatus uint8 = 2  // 等待建立完畢
	InvalidatedStatus             uint8 = 99 // 已失效
)

type Worker struct {
	ID            *uint64    `json:"id"`
	Type          *string    `form:"type" json:"type" sendForm:"type"`
	URL           *string    `form:"url" json:"url" sendForm:"url"`
	Account       *string    `form:"account" json:"account" sendForm:"account"`
	Name          *string    `form:"name" json:"name" sendForm:"name"`
	Phone         *string    `form:"phone" json:"phone" sendForm:"phone"`
	PostalCode    *string    `form:"postalCode" json:"postalCode" sendForm:"postalCode"`
	Birthday      *time.Time `form:"birthday" json:"birthday" sendForm:"birthday"`
	Email         *string    `form:"email" json:"email" sendForm:"email"`
	SimCardNumber *string    `form:"simCardNumber" json:"simCardNumber" sendForm:"simCardNumber"`
	ActivationAt  *time.Time `form:"activationAt" json:"activationAt" sendForm:"activationAt"`
	Remark        *string    `form:"remark" json:"remark" sendForm:"remark"`
	Status        *uint8     `form:"status" json:"status" sendForm:"status"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
}

type GetWorkersRes struct {
	Workers []Worker `json:"workers"`
	Count   int      `json:"count"`
}
