package core

import (
	"time"
)

type Condition struct {
	PreProcess  func() (bool, error)
	Condition   func() (bool, error)
	PostProcess func() (bool, error)
}

type Content struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	Cid        string `json:"cid"`
	BucketUuid string `json:"bucket_uuid"`
	Created_at time.Time
	Updated_at time.Time
}

type ContentDeals struct {
	ID          uint   `gorm:"primary_key"`
	ContentID   uint   `gorm:"content_id"`
	DealID      uint   `gorm:"deal_id"`
	Status      string `gorm:"status"` // active, inactive.
	Replication int    `gorm:"replication"`
	Created_at  time.Time
	Updated_at  time.Time
}

//func CreateContent(content Content) (Content, error) {
//
//	if err != nil {
//		return content, err
//	}
//	database.Create(&content)
//	return content, nil
//}
