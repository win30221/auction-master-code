package bonuses

import "time"

type Bonus struct {
	ConsignorID *uint64    `form:"consignorID" json:"consignorID"`
	Balance     *float64   `form:"balance" json:"balance"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

type BonusLog struct {
	ID              *uint64    `json:"id"`
	ConsignorID     *uint64    `form:"consignorID" json:"consignorID"`
	OpCode          *string    `form:"opCode" json:"opCode"`
	Action          *uint16    `form:"action" json:"action"`
	PreviousBalance *float64   `form:"previousBalance" json:"previousBalance"`
	NetDifference   *float64   `form:"netDifference" json:"netDifference"`
	CreatedAt       *time.Time `json:"createdAt"`
}

type GetBonusLogsReq struct {
	ConsignorID uint64    `form:"consignorID"` // 如果改成多個，做快取會有刪不到的風險
	Action      []uint16  `form:"action"`
	StartAt     time.Time `form:"startAt"`
	EndAt       time.Time `form:"endAt"`
	Sort        []string  `form:"sort"`
	Order       []string  `form:"order"`
	Limit       int64     `form:"limit"`
	Offset      int64     `form:"offset"`
}

type GetBonusLogsRes struct {
	BonusLogs []BonusLog `json:"bonusLogs"`
	Count     int64      `json:"count"`
}
