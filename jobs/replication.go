package jobs

import (
	"context"
	"fmt"
	"github.com/filecoin-project/boost/transport/httptransport"
	"github.com/filecoin-project/go-address"
	cborutil "github.com/filecoin-project/go-cbor-util"
	"github.com/filecoin-project/go-fil-markets/storagemarket/network"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multiaddr"
	"golang.org/x/xerrors"
	"light-estuary-node/core"
)

// ReplicationProcessor check replication if exists
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
			err := r.makeStorageDeal(&bucketReplicationRequest)
			if err != nil {
				fmt.Println(err)
				return
			}

			bucketReplicationRequest.Status = "complete"
			r.LightNode.DB.Save(&bucketReplicationRequest)

		}
	}

}

func (r *ReplicationProcessor) makeStorageDeal(bucketReplicationRequests *core.BucketReplicationRequest) error {

	var pieceCommitment core.PieceCommitment
	r.LightNode.DB.Model(&core.PieceCommitment{}).Where("piece = ?", bucketReplicationRequests.PieceCommitment).Find(&pieceCommitment)

	// 6 deals
	fmt.Println("pieceCommitment.Cid", pieceCommitment.Cid)
	pCid, err := cid.Decode(pieceCommitment.Cid)
	if err != nil {

	}
	fmt.Println("pCid", pCid.String())
	priceBigInt, err := types.BigFromString("0001")
	var DealDuration = 1555200 - (2880 * 21)
	duration := abi.ChainEpoch(DealDuration)

	fmt.Println(r.GetStorageProviders()[0].Address)
	prop, err := r.LightNode.Filclient.MakeDeal(context.Background(), r.GetStorageProviders()[0].Address, pCid, priceBigInt, abi.PaddedPieceSize(pieceCommitment.PaddedPieceSize), duration, true)
	fmt.Println(prop)
	if err != nil {
		return err
	}
	//fmt.Println(prop.DealProposal.Proposal.PieceCID)
	propnd, err := cborutil.AsIpld(prop.DealProposal)
	//propnd, err := cborutil.AsIpld(prop.DealProposal)
	if err != nil {
		return err
	}

	dealUUID := uuid.New()
	//deal := &model.ContentDeal{

	//deal := &model.ContentDeal{
	//	Content:             content.ID,
	//	PropCid:             util.DbCID{CID: propnd.Cid()},
	//	DealUUID:            dealUUID.String(),
	//	Miner:               miner.String(),
	//	Verified:            m.cfg.Deal.IsVerified,
	//	UserID:              content.UserID,
	//	DealProtocolVersion: proto,
	//	MinerVersion:        ask.MinerVersion,
	//}

	//if err := r.LightNode.DB.Create(nil).Error; err != nil {
	//	return xerrors.Errorf("failed to create database entry for deal: %w", err)
	//}
	fmt.Println(r.GetStorageProviders()[0].Address)
	propPhase, err := r.sendProposalV120(context.Background(), *prop, propnd.Cid(), dealUUID, 0)

	if err != nil {
		fmt.Println(err)
		//return err
	}

	fmt.Println(propPhase)
	proto, err := r.LightNode.Filclient.DealProtocolForMiner(context.Background(), r.GetStorageProviders()[0].Address)
	if err != nil {
		fmt.Println(err)
		//return err
	}
	fmt.Println(proto)
	//isPushTransfer := proto == filclient.DealProtocolv110
	// set the piece commitment status to complete
	pieceCommitment.Status = "complete"
	r.LightNode.DB.Save(&pieceCommitment)

	// set the bucket status to complete
	var bucket core.Bucket
	r.LightNode.DB.Model(&core.Bucket{}).Where("uuid = ?", bucketReplicationRequests.BucketUuid).Find(&bucket)
	bucket.Status = "complete"
	r.LightNode.DB.Save(&bucket)

	return nil

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
			fmt.Println("error on miner address", err, a)
		}
		storageProviders = append(storageProviders, MinerAddress{Address: a})
	}
	return storageProviders
}

var mainnetMinerStrs = []string{
	"f01907556",
}

func (r *ReplicationProcessor) sendProposalV120(ctx context.Context, netprop network.Proposal, propCid cid.Cid, dealUUID uuid.UUID, dbid uint) (bool, error) {
	// In deal protocol v120 the transfer will be initiated by the
	// storage provider (a pull transfer) so we need to prepare for
	// the data request

	// Create an auth token to be used in the request
	authToken, err := httptransport.GenerateAuthToken()
	if err != nil {
		return false, xerrors.Errorf("generating auth token for deal: %w", err)
	}

	rootCid := netprop.Piece.Root
	size := netprop.Piece.RawBlockSize
	var announceAddr multiaddr.Multiaddr

	if len(r.LightNode.Node.Config.AnnounceAddrs) == 0 {
		return false, xerrors.Errorf("cannot serve deal data: no announce address configured for estuary node")
	}

	addrstr := r.LightNode.Node.Config.AnnounceAddrs[0] + "/p2p/" + r.LightNode.Node.Host.ID().String()
	announceAddr, err = multiaddr.NewMultiaddr(addrstr)
	if err != nil {
		return false, xerrors.Errorf("cannot parse announce address '%s': %w", addrstr, err)
	}

	// Add an auth token for the data to the auth DB
	err = r.LightNode.Filclient.Libp2pTransferMgr.PrepareForDataRequest(ctx, dbid, authToken, propCid, rootCid, size)
	if err != nil {
		return false, xerrors.Errorf("preparing for data request: %w", err)
	}

	// Send the deal proposal to the storage provider
	propPhase, err := r.LightNode.Filclient.SendProposalV120(ctx, dbid, netprop, dealUUID, announceAddr, authToken)
	return propPhase, err
}
