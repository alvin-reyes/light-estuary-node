package core

import (
	"github.com/filecoin-project/go-state-types/abi"
)

const TestSectorSize abi.SectorSize = 512 << 20

// https://github.com/application-research/filclient-unstable/blob/73c6aa077bb3d92ecae9958280b6df4cc70cac10/common_test.go#L33
//func CreateExportFile() {
//	client, miner, _, fc, closer := initEnsemble(t, ctx)
//	importRes := CreateDummyDeal(ctx, t, client, miner)
//
//	// Transfer dummy deal into the client
//	fmt.Printf("Transferring...\n")
//	transfer, err := fc.StorageProviderByAddress(miner.ActorAddr).StartRetrievalTransfer(ctx, importRes.Root)
//	require.NoError(t, err)
//	<-transfer.Done()
//	fmt.Printf("Finished transferring\n")
//
//	// Export to file and check that it exists
//	require.NoError(t, fc.ExportToFile(ctx, importRes.Root, outputFilename, false))
//	outFile, err := os.Stat(outputFilename)
//
//	require.NoError(t, err)
//	require.Greater(t, outFile.Size(), int64(0))
//}
//func CreateDummyDeal(ctx context.Context, t *testing.T, client *kit.TestFullNode, miner *kit.TestMiner) *api.ImportRes {
//	// Create dummy deal on miner
//	res, file := client.CreateImportFile(ctx, 1, int(TestSectorSize/2))
//	fmt.Printf("Created import file '%s'\n", file)
//	pieceInfo, err := client.ClientDealPieceCID(ctx, res.Root)
//	require.NoError(t, err)
//	dh := kit.NewDealHarness(t, client, miner, miner)
//	dp := dh.DefaultStartDealParams()
//	dp.EpochPrice.Set(big.NewInt(250_000_000))
//	dp.DealStartEpoch = abi.ChainEpoch(4 << 10)
//	dp.Data = &storagemarket.DataRef{
//		TransferType: storagemarket.TTManual,
//		Root:         res.Root,
//		PieceCid:     &pieceInfo.PieceCID,
//		PieceSize:    pieceInfo.PieceSize.Unpadded(),
//	}
//	proposalCid := dh.StartDeal(ctx, dp)
//
//	carFileDir := t.TempDir()
//	carFilePath := filepath.Join(carFileDir, "out.car")
//	fmt.Printf("Generating car...\n")
//	require.NoError(t, client.ClientGenCar(ctx, api.FileRef{Path: file}, carFilePath))
//	fmt.Printf("Importing car...\n")
//	require.NoError(t, miner.DealsImportData(ctx, *proposalCid, carFilePath))
//
//	dh.StartSealingWaiting(ctx)
//	dh.WaitDealPublished(ctx, proposalCid)
//
//	return res
//}
