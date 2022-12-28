package core

import "fmt"

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
