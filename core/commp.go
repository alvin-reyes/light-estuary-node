package core

import "fmt"

func ComputeCommp(bucketUui string) {

	database, err := OpenDatabase()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database opened", database)
	// get content with bucket UUI and create a CAR file

	// save the cid, commp and bucket UUI to the database.
}
