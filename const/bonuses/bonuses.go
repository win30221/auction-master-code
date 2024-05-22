package bonuses

import "time"

type Bonus struct {
	ConsignorID *uint64    `form:"consignorID" json:"consignorID"`
	Balance     *float64   `form:"balance" json:"balance"`
	CreatedAt   *time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt   *time.Time `form:"updatedAt" json:"updatedAt"`
}

type BonusLog struct {
	ID              *uint64    `json:"id"`
	ConsignorID     *uint64    `form:"consignorID" json:"consignorID"`
	OpCode          *string    `form:"opCode" json:"opCode"`
	Action          *string    `form:"action" json:"action"`
	PreviousBalance *float64   `form:"previousBalance" json:"previousBalance"`
	NetDifference   *float64   `form:"netDifference" json:"netDifference"`
	CreatedAt       *time.Time `form:"createdAt" json:"createdAt"`
}
