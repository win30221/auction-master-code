package workerhubs

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// Worker 類型
	SellerType  = "Seller"
	WatcherType = "Watcher"

	// Worker 狀態
	ActiveStatus                  uint8 = 1  // 啟用中
	AwaitingSetupCompletionStatus uint8 = 2  // 等待建立完畢
	InvalidatedStatus             uint8 = 99 // 已失效
)

type Worker struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id" form:"id"`
	Type          string             `bson:"type" json:"type" form:"type"`
	URL           string             `bson:"url" json:"url" form:"url"`
	Status        uint8              `bson:"status" json:"status" form:"status"`
	Location      string             `bson:"location" json:"location" form:"location"`
	Name          string             `bson:"name" json:"name" form:"name"`
	Phone         string             `bson:"phone" json:"phone" form:"phone"`
	PostalCode    string             `bson:"postalCode" json:"postalCode" form:"postalCode"`
	Birthday      time.Time          `bson:"birthday" json:"birthday" form:"birthday"`
	Email         string             `bson:"email" json:"email" form:"email"`
	SimCardNumber string             `bson:"simCardNumber" json:"simCardNumber" form:"simCardNumber"`
	ActivationAt  time.Time          `bson:"activationAt" json:"activationAt" form:"activationAt"`
	Remark        string             `bson:"remark" json:"remark" form:"remark"`
}
