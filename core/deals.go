package core

import (
	"github.com/filecoin-project/go-state-types/abi"
	"time"
)

const TestSectorSize abi.SectorSize = 512 << 20

type Deals struct {
	ID         int    `gorm:"primaryKey"`
	DealID     int    `gorm:"deal_id"`
	BucketUuid string `gorm:"bucket_uuid"`
	Created_at time.Time
	Updated_at time.Time
}
