package jobs

import (
	"context"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-merkledag"
	"light-estuary-node/core"
	"time"
)

// workers
// jobs

// this processors are independent. we want it to run on it's own without waiting
// for other groups.

type CarGeneratorProcessor struct {
	Processor
}

func NewCarGeneratorProcessor(ln *core.LightNode) CarGeneratorProcessor {
	return CarGeneratorProcessor{
		Processor{
			LightNode: ln,
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
		r.LightNode.DB.Model(&core.Content{}).Where("bucket_uuid = ?", bucket.UUID).Find(&contents)
		rootCid, err := r.buildCarForListOfContents(contents)
		if err != nil {
			panic(err)
		}

		// update bucket cid and status
		//r.LightNode.DB.Model(&core.Bucket{}).x
		bucket.Updated_at = time.Now()
		bucket.Cid = rootCid.String()
		bucket.Status = "car-assigned"

		r.LightNode.DB.Updates(&bucket)

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
		//fmt.Println("Adding node: ", rootCid.String())
		r.addToBlockstore(r.LightNode.Node.DAGService, &node)
	}

	// add to blockstore

	return rootCid, nil
}

func (r *CarGeneratorProcessor) addToBlockstore(ds format.DAGService, nds ...format.Node) {
	for _, nd := range nds {
		//fmt.Println("Adding node: ", nd.Cid().String())
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
