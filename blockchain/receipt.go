package blockchain

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// GetGasReceipt extracts and returns gas details from the transaction
// receipt of the given tx hash
func GetGasReceipt(ec *ethclient.Client, txHash common.Hash) *GasReceipt {

	result := &GasReceipt{
		TxHash:            txHash,
		GasUsed:           0,
		CumulativeGasUsed: 0,
	}

	rcpt, err := ec.TransactionReceipt(context.Background(), txHash)

	if err != nil {
		return result //Return trivial values
	}

	result.GasUsed = rcpt.GasUsed
	result.CumulativeGasUsed = rcpt.CumulativeGasUsed
	return result

}

// GetReceiptLite extracts and returns (relevant to Pontus) details from
// the transaction receipt of the given txhash
func GetReceiptLite(c *rpc.Client, txHash string) (*ReceiptLite, error) {

	var result *ReceiptLite

	err := c.Call(&result, "eth_getTransactionReceipt", common.HexToHash(txHash))
	if err != nil {
		return nil, err
	}

	return result, nil
}

// ReceiptLite represents a subset of a transaction receipt
// to store receipt data relevant to SKALE
type ReceiptLite struct {
	TxHash string       `json:"transactionHash" gencodec:"required"`
	Logs   []*types.Log `json:"logs"              gencodec:"required"`
	From   string       `json:"from"`
	To     string       `json:"to"`
}

// GasReceipt is a struct to store gas related details from a
// transaction receipt
type GasReceipt struct {
	TxHash            common.Hash `json:"transactionHash"`
	CumulativeGasUsed uint64      `json:"cumulativeGasUsed"`
	GasUsed           uint64      `json:"gasUsed"`
}
