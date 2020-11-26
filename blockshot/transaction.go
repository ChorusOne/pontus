package blockshot

import (
	"strconv"

	"github.com/ChorusOne/pontus-internal/blockchain"
	"github.com/ChorusOne/pontus-internal/connection"
	"github.com/ethereum/go-ethereum/core/types"

	"math/big"
)

//ProcessTransactions is a task to retrieve all transactions of the given block number,
//extract relevant information for each transaction and store it in the database
func ProcessTransactions(blockNumber *big.Int) error {

	b, err := blockchain.GetBlock(blockNumber)

	if err != nil {
		return err
	}

	transactions := b.Transactions()
	timestamp := strconv.FormatUint(b.Time(), 10)

	for _, tx := range transactions {
		gasReceipt := blockchain.GetGasReceipt(connection.ETHCLIENT, tx.Hash())
		processTransaction(blockNumber, timestamp, tx, gasReceipt)
	}

	return nil

}

//processTransaction parses and dumps to database - the information of a single transaction
func processTransaction(blockNumber *big.Int, timestamp string, tx *types.Transaction, gasReceipt *blockchain.GasReceipt) error {

	transaction := TransactionFrom(tx, gasReceipt)
	DumpTransaction(blockNumber, timestamp, tx.Hash().String(), transaction)

	return nil

}
