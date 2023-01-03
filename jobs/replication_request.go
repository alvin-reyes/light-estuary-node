package jobs

import (
	"light-estuary-node/core"
	"time"
)

type ReplicationRequestProcessor struct {
	Processor
}

func NewReplicationRequestProcessor(ln *core.LightNode) ReplicationRequestProcessor {
	return ReplicationRequestProcessor{
		Processor{
			LightNode: ln,
		},
	}
}

func (r *ReplicationRequestProcessor) Run() {

	// for all piece commitment records, check if status is open
	var pieceCommitments []core.PieceCommitment
	r.LightNode.DB.Model(&core.PieceCommitment{}).Where("status = ?", "open").Find(&pieceCommitments)

	// for each piece commitment, create replication request
	for _, pieceCommitment := range pieceCommitments {
		for i := 0; i < 6; i++ {
			r.LightNode.DB.Create(&core.BucketReplicationRequest{
				BucketUuid:      pieceCommitment.BucketUuid,
				PieceCommitment: pieceCommitment.Piece,
				Cid:             pieceCommitment.Cid,
				Status:          "open",
				Created_at:      time.Now(),
			})
		}

		// replication assigned
		pieceCommitment.Status = "replication-assigned"
		r.LightNode.DB.Save(&pieceCommitment)
	}

}
