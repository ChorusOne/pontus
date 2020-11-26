package atlas

import (
	"log"
	"math/big"
	"strings"

	"github.com/ChorusOne/pontus-internal/connection"

	skaletoken "github.com/ChorusOne/pontus-internal/binding"
	"github.com/ChorusOne/pontus-internal/constants"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// GetBlockTransferAddresses extracts addresses from a SKALE transfer event logs.
// TODO: modularize log slice, constants
func GetBlockTransferAddresses(fromBlock *big.Int, toBlock *big.Int) map[string]bool {
	empty := map[string]bool{}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	topics := [][]common.Hash{{logTransferSigHash}}
	log.Println("Transfer hex:")
	log.Println(logTransferSigHash.Hex())

	contractAbi, err := abi.JSON(strings.NewReader(string(skaletoken.SkaleTokenABI)))

	if err != nil {
		log.Printf("Could not load contract ABI: %s\n", err)
		return empty
	}

	logs, err := connection.GetBlockLogs(
		fromBlock,
		toBlock,
		topics,
		constants.ContractDeploymentAddress[constants.Mainnet][constants.SkaleToken])

	if err != nil {
		log.Printf("Could not read block logs: %s\n", err)
		return empty
	}

	addresses := map[string]bool{}

	for _, entry := range logs {
		if entry.Topics[0].Hex() == logTransferSigHash.Hex() {
			var transferEvent constants.LogTransfer

			err := contractAbi.Unpack(&transferEvent, "Transfer", entry.Data)

			if err != nil {
				log.Fatal(err)
			}

			transferEvent.From = common.HexToAddress(entry.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(entry.Topics[2].Hex())
			sender := transferEvent.From.Hex()
			recipient := transferEvent.To.Hex()
			addresses[sender] = true
			addresses[recipient] = true

		}
	}
	return addresses
}
