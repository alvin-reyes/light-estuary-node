package jobs

import (
	"context"
	"github.com/google/uuid"
	"light-estuary-node/core"
	"time"
)

var workerPool = make(chan struct{}, 10)

type BucketAssignProcessor struct {
	Processor
}

func NewBucketAssignProcessor() BucketAssignProcessor {
	node, err := core.NewLightNode(context.Background()) // new light node
	if err != nil {
		panic(err)
	}
	return BucketAssignProcessor{
		Processor{
			LightNode: node,
		},
	}
}

func (r *BucketAssignProcessor) Run() {
	// run the content processor.
	var contents []core.Content
	r.LightNode.DB.Model(&core.Content{}).Where("status = ? and bucket_uuid is null", "open").Find(&contents)

	// get range of content ids and assign a bucket
	// if there are contents, create a new bucket and assign it to the contents
	uuid, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	if len(contents) > 0 {
		// create a new bucket
		bucket := core.Bucket{
			Status:     "open",        // open, in-progress, completed (closed).
			Name:       uuid.String(), // same as uuid
			UUID:       uuid.String(),
			Created_at: time.Now(), // log it.
		}
		r.LightNode.DB.Create(&bucket)

		// assign bucket to contents
		r.LightNode.DB.Model(&core.Content{}).Where("status = ? and bucket_uuid is null", "open").Update("bucket_uuid", bucket.UUID)
	}

}
