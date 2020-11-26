package connection

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var ETHCLIENT *ethclient.Client
var RPCCLIENT *rpc.Client

// InitEthClient creates a new ethclient pointer for the Ethereum blockchain
func InitEthClients() {

	var IPCFilePath = os.Getenv("IPC_FILE_PATH")
	var RPCURLPath = os.Getenv("RPC_URL_PATH")

	// First try IPC because it's faster
	ec, err := ethclient.Dial(IPCFilePath)

	if err != nil {
		// Then try RPC if we can't make an IPC connection
		log.Println("Warning, not able to connect via IPC. Trying RPC but snapshotting will be slow.")

		ec, err = ethclient.Dial(RPCURLPath)
		// We can't connect to either, exit.
		if err != nil {
			log.Fatalln("Check IPC/RPC config paths to Ethereum full node.")
		}
	}

	ETHCLIENT = ec

	//First try IPC node
	//TODO: see if we need RPC since ethclient supports most queries
	rc, err := rpc.Dial(IPCFilePath)

	if err != nil {
		// Try RPC.
		log.Println("Warning, not able to connect via IPC. Trying RPC but snapshotting will be slow.")
		rc, err = rpc.Dial(RPCURLPath)
		if err != nil {
			log.Fatalln("Check IPC/RPC config paths to Ethereum full node.")
		}
	}

	RPCCLIENT = rc
}

// GetBlockLogs returns a log dump at a given Ethereum block number
func GetBlockLogs(fromBlock *big.Int, toBlock *big.Int, topics [][]common.Hash, contractAddress string) ([]types.Log, error) {

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Topics:    topics,
		Addresses: []common.Address{
			common.HexToAddress(contractAddress),
		},
	}
	logs, err := ETHCLIENT.FilterLogs(context.Background(), query)

	return logs, err
}
