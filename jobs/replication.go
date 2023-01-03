package jobs

import (
	"context"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"light-estuary-node/core"
)

// check replication if exists
type ReplicationProcessor struct {
	Processor
}

func NewReplicationProcessor(ln *core.LightNode) ReplicationProcessor {
	return ReplicationProcessor{
		Processor{
			LightNode: ln,
		},
	}
}

func (r *ReplicationProcessor) Run() {

	// get all piece comm record that are replication-assigned
	var pieceCommitments []core.PieceCommitment
	r.LightNode.DB.Model(&core.PieceCommitment{}).Where("status = ?", "replication-assigned").Find(&pieceCommitments)

	// for each piece commitment, look for bucket replication requests
	for _, pieceCommitment := range pieceCommitments {
		var bucketReplicationRequests []core.BucketReplicationRequest
		r.LightNode.DB.Model(&core.BucketReplicationRequest{}).Where("id = ?", pieceCommitment.ID).Find(&bucketReplicationRequests)

		// for each replication request, make a deal
		for _, bucketReplicationRequest := range bucketReplicationRequests {
			r.makeStorageDeal(&bucketReplicationRequest)
		}
	}

}

func (r *ReplicationProcessor) makeStorageDeal(bucketReplicationRequests *core.BucketReplicationRequest) {

	var pieceCommitment core.PieceCommitment
	r.LightNode.DB.Model(&core.PieceCommitment{}).Where("id = ?", bucketReplicationRequests.PieceCommitment).Find(&pieceCommitment)

	// 6 deals
	pCid, err := cid.Decode(pieceCommitment.Cid)
	if err != nil {

	}
	fmt.Println(pCid)
	priceBigInt, err := types.BigFromString("0001")
	var DealDuration = 1555200 - (2880 * 21)
	duration := abi.ChainEpoch(DealDuration)

	fmt.Println(r.GetStorageProviders()[0].Address)
	proposal, err := r.LightNode.Filclient.MakeDeal(context.Background(), r.GetStorageProviders()[0].Address, pCid, priceBigInt, abi.PaddedPieceSize(pieceCommitment.PaddedPieceSize), duration, true)
	fmt.Println(proposal)
	if err != nil {
		fmt.Println("err>>>>", err)
		return // don't complete but log the error on the request table
	}
	fmt.Println(proposal)

	// create the deal record
	//r.LightNode.DB.Create(&core.Deal{
	//	DealId:                  proposal.ProposalCid.String(),
	//	BucketReplicationRequest: bucketReplicationRequests.ID,
	//	Status:                  "open",
	//}

	// set the piece commitment status to complete
	pieceCommitment.Status = "complete"
	r.LightNode.DB.Save(&pieceCommitment)

	// set the bucket status to complete
	var bucket core.Bucket
	r.LightNode.DB.Model(&core.Bucket{}).Where("uuid = ?", bucketReplicationRequests.BucketUuid).Find(&bucket)
	bucket.Status = "complete"
	r.LightNode.DB.Save(&bucket)

}

type MinerAddress struct {
	Address address.Address
}

func (r *ReplicationProcessor) GetStorageProviders() []MinerAddress {
	var storageProviders []MinerAddress
	for _, s := range mainnetMinerStrs {
		address.CurrentNetwork = address.Mainnet
		a, err := address.NewFromString(s)
		if err != nil {
			panic(err)
		}
		fmt.Println("a", a)
		storageProviders = append(storageProviders, MinerAddress{Address: a})
	}
	fmt.Println(storageProviders)
	return storageProviders
}

var mainnetMinerStrs = []string{
	"f0123261",
}
