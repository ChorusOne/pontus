package blockchain

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"math/big"
)

//GetBlock returns the full block (including transactional data) for the given block number
func getBlock(ec *ethclient.Client, blockNumber *big.Int) (*types.Block, error) {

	var result *types.Block
	result, err := ec.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return nil, err
	}

	return result, nil
}
