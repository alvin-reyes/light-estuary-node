package core

import "time"

type StagingBucket struct {
	// ID is the unique identifier for the bucket.
	ID         string `json:"id"`
	Name       string `json:"name"`
	UUID       string `json:"uuid"`
	Status     string `json:"status"` // open, in-progress, completed (closed).
	Cid        string `json:"cid"`    // car file of the consolidated content
	created_at time.Time
	updated_at time.Time
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
	created_at  time.Time
	updated_at  time.Time
}
