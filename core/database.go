package core

import (
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

func OpenDatabase() (*gorm.DB, error) {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	dbName, okHost := viper.Get("DB_NAME").(string)
	if !okHost {
		panic("DB_NAME not set")
	}
	DB, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	// generate new models.
	ConfigureModels(DB) // create models.

	if err != nil {
		return nil, err
	}
	return DB, nil
}

func ConfigureModels(db *gorm.DB) {
	db.AutoMigrate(&Content{}, &Bucket{}, &StorageProvider{}, &PieceCommitment{}, &Deals{}, &BucketReplicationRequest{})
}

// replication request
type BucketReplicationRequest struct {
	ID              uint   `gorm:"primaryKey"`
	BucketUuid      string `gorm:"bucket_uuid"`
	PieceCommitment string `gorm:"piece_commitment"`
	Cid             string `gorm:"cid"`
	Status          string `gorm:"status"`
	Created_at      time.Time
	Updated_at      time.Time
}

type ContentReplicationRequest struct {
	ID              uint   `gorm:"primaryKey"`
	ContentId       string `gorm:"content_id"`
	PieceCommitment string `gorm:"piece_commitment"`
	Cid             string `gorm:"cid"`
	Status          string `gorm:"status"`
	Created_at      time.Time
	Updated_at      time.Time
}
