package core

import (
	"fmt"
	"time"
)

type PieceCommitment struct {
	ID              uint   `gorm:"primaryKey"`
	Cid             string `json:"cid"`
	Piece           string `json:"piece"`
	Size            int64  `json:"size"`
	CarSize         int64  `json:"car_size"`
	PaddedPieceSize uint64 `json:"padded_piece_size"`
	BucketUuid      string `json:"bucket_uuid"`
	Status          string `json:"status"` // open, in-progress, completed (closed).
	Created_at      time.Time
	Updated_at      time.Time
}

func CreateCarAndComputeCommp(bucketUui string) {

	database, err := OpenDatabase()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database opened", database)
	// get content with bucket UUI and create a CAR file
	database.Model(&Content{}).Where("bucket = ?", bucketUui).Find(&Content{})

	// create the CAR file

	// update content after getting the car file.

	// save the cid, commp and bucket UUI to the database.

}
