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

	// get open buckets and create a car for each content cid
	var buckets []core.Bucket
	r.LightNode.DB.Model(&core.Bucket{}).Where("status = ?", "open").Find(&buckets)

	//	for each bucket, get the contents and create a car
	for _, bucket := range buckets {
		var contents []core.Content
		r.LightNode.DB.Model(&core.Content{}).Where("status = ? and bucket_uuid = ?", "open", bucket.UUID).Find(&contents)
		rootCid, err := r.buildCarForListOfContents(contents)
		if err != nil {
			panic(err)
		}
		// update the bucket
		r.LightNode.DB.Model(&core.Bucket{}).Where("uuid = ?", bucket.UUID).Update("cid", rootCid.String())
		// update contents.
		r.LightNode.DB.Model(&core.Content{}).Where("status = ? and bucket_uuid = ?", "open", bucket.UUID).Update("status", "assigned")
	}
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
		r.addToBlockstore(r.LightNode.Node.DAGService, &node)
	}

	// add to blockstore

	return rootCid, nil
}

func (r *CarGeneratorProcessor) addToBlockstore(ds format.DAGService, nds ...format.Node) {
	for _, nd := range nds {
		fmt.Println("Adding node: ", nd.Cid().String())
		if err := ds.Add(context.Background(), nd); err != nil {
			panic(err)
		}
	}
}

func (r *CarGeneratorProcessor) getNodeForCid(content core.Content) (format.Node, error) {
	decodedCid, err := cid.Decode(content.Cid)
	if err != nil {
		return nil, err
	}
	return r.LightNode.Node.Get(context.Background(), decodedCid)
}
