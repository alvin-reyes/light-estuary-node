package core

import (
	"time"
)

type Condition struct {
	PreProcess  func() (bool, error)
	Condition   func() (bool, error)
	PostProcess func() (bool, error)
}

type PieceCommitment struct {
	ID         uint   `gorm:"primaryKey"`
	Cid        string `json:"cid"`
	Piece      string `json:"piece"`
	Size       int64  `json:"size"`
	CarSize    int64  `json:"car_size"`
	Created_at time.Time
	Updated_at time.Time
}
type Content struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	Cid           string `json:"cid"`
	StagingBucket string `json:"staging_bucket"` // where this content will be associated
	Created_at    time.Time
	Updated_at    time.Time
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
