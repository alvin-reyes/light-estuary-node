package jobs

import (
	"fmt"
	"light-estuary-node/core"
)

type DealsProcessor struct {
	Processor
}

func NewDealsProcessor() DealsProcessor {
	return DealsProcessor{}
}

func (r *DealsProcessor) Run() {

	// for all piece commitment records, check if status is open

	var pieceCommitments []core.PieceCommitment
	r.LightNode.DB.Model(&core.PieceCommitment{}).Where("status = ?", "open").Find(&pieceCommitments)

	for _, pieceCommitment := range pieceCommitments {

		// update status of piece commitment to in-progress
		fmt.Println(pieceCommitment.BucketUuid)

	}

}

func (r *DealsProcessor) makeStorageDeal() {

	// 6 deals

}

func (r *DealsProcessor) getMiners() {

}
