package jobs

import (
	"context"
	"github.com/application-research/filclient"
	"github.com/ipfs/go-cid"
	"light-estuary-node/core"
	"time"
)

// workers
// jobs

// this processors are independent. we want it to run on it's own without waiting
// for other groups.

type CommpProcessor struct {
	Processor
}

func NewCommpProcessor(ln *core.LightNode) CommpProcessor {
	return CommpProcessor{
		Processor{
			LightNode: ln,
		},
	}
}

func (r *CommpProcessor) Run() {

	// get the CID field of the bucket and generate a commp for it.
	var buckets []core.Bucket
	r.LightNode.DB.Model(&core.Bucket{}).Where("status = ?", "car-assigned").Find(&buckets)

	// for each bucket, generate a commp
	for _, bucket := range buckets {
		// update the bucket
		r.LightNode.DB.Model(&core.Bucket{}).Where("uuid = ?", bucket.UUID).Update("status", "in-progress")
		payloadCid, err := cid.Decode(bucket.Cid)
		if err != nil {
			panic(err)
		}

		//return commCid, preparedCar.Size(), abi.PaddedPieceSize(size).Unpadded(), nil
		commitment, u, a, err := filclient.GeneratePieceCommitment(context.Background(), payloadCid, r.LightNode.Node.Blockstore)
		if err != nil {
			return
		}

		// save the commp to the database
		r.LightNode.DB.Create(&core.PieceCommitment{
			Cid:             payloadCid.String(),
			Piece:           commitment.String(),
			Size:            int64(u),
			PaddedPieceSize: uint64(a),
			Status:          "open",
			BucketUuid:      bucket.UUID,
			Created_at:      time.Now(),
			Updated_at:      time.Now(),
		})

		// update bucket status to commp-computed
		r.LightNode.DB.Model(&core.Bucket{}).Where("uuid = ?", bucket.UUID).Update("status", "piece-assigned")
	}
}
