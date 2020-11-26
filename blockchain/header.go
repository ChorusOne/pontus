package blockchain

import (
	"context"
	"log"
	"math/big"

	"github.com/ChorusOne/pontus-internal/constants"

	"github.com/ChorusOne/pontus-internal/connection"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

// GetHeader returns the header of an Ethereum block at a given block number
func GetHeader(blockNumber *big.Int) *types.Header {

	h, err := connection.ETHCLIENT.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatalf("Fatal error, could not retrieve ETH block header: %s at block number: %s ", err, blockNumber)
	}

	return h
}

// GetLatestBlockNumber returns the block at the current height of the Ethereum blockchain wrt SKALE genesis
func GetLatestBlockNumber() *big.Int {
	var number string

	err := connection.RPCCLIENT.Call(&number, "eth_blockNumber")
	if err != nil {
		log.Printf("Unable to retrieve latest block number: %s, returning SKALE genesis\n", err)
		return constants.GenesisBlockNumber[constants.NetActive]
	}
	// JSON RPC returns latest block number as a hex string, convert to int
	bn, err := hexutil.DecodeBig(number)
	if err != nil {
		log.Printf("Unable to decode latest block number: %s", err)
	}

	return bn
}
