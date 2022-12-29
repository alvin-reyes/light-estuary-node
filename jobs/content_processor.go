package jobs

import (
	"context"
	"light-estuary-node/core"
)

var workerPool = make(chan struct{}, 10)

type ContentProcessor struct {
	Node *core.LightNode
}

func NewContentProcessor() ContentProcessor {
	node, err := core.NewLightNode(context.Background()) // new light node
	if err != nil {
		panic(err)
	}
	return ContentProcessor{
		Node: node,
	}
}

func (r *ContentProcessor) Run() {
	// run the content processor.
	r.Node.DB.Model(&core.Content{}).Where("status = ? and bucket is null", "open").Find(&core.Content{})

	// get collection of files and compute size (if it's more than 1GB) assign it.

	// if it's time, get the files and just assign to a new bucket.

	// create a bucket for tracking and set it to open.
}
