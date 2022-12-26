package jobs

type Condition struct {
	PreProcess  func() (bool, error)
	Condition   func() (bool, error)
	PostProcess func() (bool, error)
}
type BucketProcessor struct {
	Bucket        string
	JobConditions Condition // we need to satisfy all conditions.
}

// checks the buckets for new files and compare them against conditions
func RunProcessor() {
	// 1. get all the buckets
	// 2. for each bucket, get all the files
	// 3. for each file, check against conditions
	// 4. if all conditions are satisfied, then process the file
}

func BuildCarFileFromBucket() {

}

func ComputeCommitmentPiece() {

}

func GetAllPossibleMiners() {

}
