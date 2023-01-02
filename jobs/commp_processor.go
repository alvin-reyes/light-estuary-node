package jobs

import (
	"context"
	"light-estuary-node/core"
)

// workers
// jobs

// this processors are independent. we want it to run on it's own without waiting
// for other groups.

type CommpProcessor struct {
	Processor
}

func NewCommpProcessor() ContentProcessor {
	node, err := core.NewLightNode(context.Background()) // new light node
	if err != nil {
		panic(err)
	}
	return ContentProcessor{
		Processor{
			LightNode: node,
		},
	}
}

func (r *CommpProcessor) Run() {

	// get the cid of the bucket

	//filclient.GeneratePieceCommitment(context.Background())

	// insert commp into table
}
