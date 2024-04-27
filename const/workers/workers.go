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
	Type          *string    `form:"type" json:"type"`
	URL           *string    `form:"url" json:"url"`
	Account       *string    `form:"account" json:"account"`
	Name          *string    `form:"name" json:"name"`
	Phone         *string    `form:"phone" json:"phone"`
	PostalCode    *string    `form:"postalCode" json:"postalCode"`
	Birthday      *time.Time `form:"birthday" json:"birthday"`
	Email         *string    `form:"email" json:"email"`
	SimCardNumber *string    `form:"simCardNumber" json:"simCardNumber"`
	ActivationAt  *time.Time `form:"activationAt" json:"activationAt"`
	Remark        *string    `form:"remark" json:"remark"`
	Status        *uint8     `form:"status" json:"status"`
	CreatedAt     *time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt     *time.Time `form:"updatedAt" json:"updatedAt"`
}
