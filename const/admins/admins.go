package admins

import "time"

var (
	// Admin 狀態
	EnabledStatus  uint8 = 1  // 啟用中
	DisabledStatus uint8 = 99 // 禁用中
)

type Admin struct {
	ID        *uint64    `json:"id"`
	Account   *string    `form:"account" json:"account"`
	Password  *string    `form:"password" json:"password"` // Argon2 encrypted
	Status    *uint8     `form:"status" json:"status"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type GetAdminsRes struct {
	Admins []Admin `json:"admins"`
	Count  int     `json:"count"`
}
