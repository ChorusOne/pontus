package blockshot

import (
	"math/big"

	"github.com/ChorusOne/pontus-internal/blockchain"
	"github.com/ChorusOne/pontus-internal/constants"
	"github.com/ChorusOne/pontus-internal/snapshot"
)

//TakeDailySnapshots evaluates if the block number is first block of UTC day
//and then takes a snapshot if that is indeed the case.
func TakeDailySnapshots(blockNumber *big.Int) error {

	if isFirstBlockOfDay(blockNumber) {

		if isGenesisBlock(blockNumber) {
			snapshot.GenerateAndStoreGenesisSnapshot(blockNumber)
		} else {
			snapshot.GenerateAndStoreBlockSnapshots(blockNumber)

		}
	}
	return nil
}

func isGenesisBlock(blockNumber *big.Int) bool {
	return blockNumber.Cmp(constants.GenesisBlockNumber[constants.NetActive]) == 0
}

func isFirstBlockOfDay(blockNumber *big.Int) bool {

	if isGenesisBlock(blockNumber) {
		return true
	}

	previousBlockHeader := blockchain.GetHeader(big.NewInt(0).Sub(blockNumber, big.NewInt(1)))
	currentBlockHeader := blockchain.GetHeader(blockNumber)

	return snapshot.SameDateOfTimestamps(previousBlockHeader.Time, currentBlockHeader.Time)
}
