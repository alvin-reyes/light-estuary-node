package jobs

import (
	"fmt"
	"light-estuary-node/core"
)

type DealsProcessor struct {
	Processor
}

func NewDealsProcessor(ln *core.LightNode) DealsProcessor {
	return DealsProcessor{
		Processor{
			LightNode: ln,
		},
	}
}

func (r *DealsProcessor) Run() {

	// for all piece commitment records, check if status is open

	var pieceCommitments []core.PieceCommitment
	r.LightNode.DB.Model(&core.PieceCommitment{}).Where("status = ?", "open").Find(&pieceCommitments)

	for _, pieceCommitment := range pieceCommitments {

		// update status of piece commitment to in-progress
		fmt.Println("making a deal for this piece")
		fmt.Println(pieceCommitment.BucketUuid)
		fmt.Println(pieceCommitment.Cid)

		// create a deal
		r.makeStorageDeal(&pieceCommitment)

		// insert deal record

		// update bucket record

	}

}

func (r *DealsProcessor) makeStorageDeal(pieceCommittment *core.PieceCommitment) {

	// 6 deals
	//r.LightNode.Filclient.MakeDeal()

}

func (r *DealsProcessor) getStorageProviders() *[]core.StorageProviders {

	return nil
}
