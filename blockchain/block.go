package blockchain

import (
	"context"
	"github.com/ChorusOne/pontus-internal/connection"
	"github.com/ethereum/go-ethereum/core/types"

	"math/big"
)

//GetBlock returns the full block (including transactional data) for the given block number
func GetBlock(blockNumber *big.Int) (*types.Block, error) {
	var result *types.Block
	result, err := connection.ETHCLIENT.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return nil, err
	}

	return result, nil
}