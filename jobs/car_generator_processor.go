package jobs

import (
	"context"
	"fmt"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-merkledag"
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

	// get the array of content that are open and have no bucket assigned.
	var contents []core.Content
	// run the content processor.
	r.LightNode.DB.Model(&core.Content{}).Where("status = ? and bucket is null", "open").Find(&contents)

	// get collection of files and compute size (if it's more than 1GB) assign it.
	rootCid, err := r.buildCarForListOfContents(contents)
	if err != nil {
		panic(err)
	}

	fmt.Println(rootCid)
	// once we have the root cid, we update the files
	// create a bucket for tracking and set it to open.
}

func (r *CarGeneratorProcessor) buildCarForListOfContents(contents []core.Content) (cid.Cid, error) {
	var rootCid cid.Cid
	for i, content := range contents {
		node := merkledag.ProtoNode{}
		nodeFromCid, err := r.getNodeForCid(content)
		if err != nil {
			return cid.Undef, err
		}
		node.AddNodeLink(nodeFromCid.String(), nodeFromCid)

		// when last index
		if i == len(contents)-1 {
			rootCid = node.Cid()
		}
	}

	return rootCid, nil
}

func (r *CarGeneratorProcessor) getNodeForCid(content core.Content) (format.Node, error) {
	decodedCid, err := cid.Decode(content.Cid)
	if err != nil {
		return nil, err
	}
	return r.LightNode.Node.Get(context.Background(), decodedCid)
}
