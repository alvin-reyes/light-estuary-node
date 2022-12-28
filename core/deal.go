package core

import (
	"context"
	"fmt"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/itests/kit"
	"github.com/stretchr/testify/require"
	"math/big"
	"path/filepath"
	"testing"
)

const TestSectorSize abi.SectorSize = 512 << 20

func CreateDummyDeal(ctx context.Context, t *testing.T, client *kit.TestFullNode, miner *kit.TestMiner) *api.ImportRes {
	// Create dummy deal on miner
	res, file := client.CreateImportFile(ctx, 1, int(TestSectorSize/2))
	fmt.Printf("Created import file '%s'\n", file)
	pieceInfo, err := client.ClientDealPieceCID(ctx, res.Root)
	require.NoError(t, err)
	dh := kit.NewDealHarness(t, client, miner, miner)
	dp := dh.DefaultStartDealParams()
	dp.EpochPrice.Set(big.NewInt(250_000_000))
	dp.DealStartEpoch = abi.ChainEpoch(4 << 10)
	dp.Data = &storagemarket.DataRef{
		TransferType: storagemarket.TTManual,
		Root:         res.Root,
		PieceCid:     &pieceInfo.PieceCID,
		PieceSize:    pieceInfo.PieceSize.Unpadded(),
	}
	proposalCid := dh.StartDeal(ctx, dp)

	carFileDir := t.TempDir()
	carFilePath := filepath.Join(carFileDir, "out.car")
	fmt.Printf("Generating car...\n")
	require.NoError(t, client.ClientGenCar(ctx, api.FileRef{Path: file}, carFilePath))
	fmt.Printf("Importing car...\n")
	require.NoError(t, miner.DealsImportData(ctx, *proposalCid, carFilePath))

	dh.StartSealingWaiting(ctx)
	dh.WaitDealPublished(ctx, proposalCid)

	return res
}
