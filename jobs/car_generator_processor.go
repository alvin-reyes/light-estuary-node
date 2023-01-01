package jobs

import (
	"context"
	"light-estuary-node/core"
)

// workers
// jobs

// this processors are independent. we want it to run on it's own without waiting
// for other groups.

type CarGeneratorProcessor struct {
	Processor
}

func NewCarGeneratorProcessor() CarGeneratorProcessor {
	node, err := core.NewLightNode(context.Background()) // new light node
	if err != nil {
		panic(err)
	}
	return CarGeneratorProcessor{
		Processor{
			LightNode: node,
		},
	}
}

func (r *CarGeneratorProcessor) Run() {

	// run the content processor.
	r.LightNode.DB.Model(&core.Content{}).Where("status = ? and bucket is null", "open").Find(&core.Content{})

	// get collection of files and compute size (if it's more than 1GB) assign it.

	// if it's time, get the files and just assign to a new bucket

	// create a bucket for tracking and set it to open.
}
