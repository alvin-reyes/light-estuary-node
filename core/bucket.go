package core

import (
	"github.com/google/uuid"
	"time"
)

// DB models
type Bucket struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `json:"name"`
	UUID       string `json:"uuid"`
	Status     string `json:"status"` // open, in-progress, completed (closed).
	Cid        string `json:"cid"`    // car file of the consolidated content
	Created_at time.Time
	Updated_at time.Time
}

type BucketManager struct {
}

func NewBucket(name string) Bucket {
	return *CreateDefaultBucket(name)
}

func CreateDefaultBucket(name string) *Bucket {
	// Create predefined buckets
	uuid, _ := uuid.NewUUID()

	bucket := &Bucket{
		Name:       name,
		UUID:       uuid.String(),
		Status:     "open",
		Cid:        "",
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}
	return bucket
}

func GetRandomBucket() error {
	return nil
}
