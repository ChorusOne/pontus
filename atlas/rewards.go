package atlas

import (
	"log"
	"math/big"
	"strings"

	"github.com/ChorusOne/pontus-internal/binding"
	"github.com/ChorusOne/pontus-internal/connection"
	"github.com/ChorusOne/pontus-internal/constants"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// BountyWasPaidEvent represents data for the event BountyWasPaid(uint,uint)"
type BountyWasPaidEvent struct {
	ValidatorID *big.Int
	Amount      *big.Int
}

// Bounty represents a SKALE bounty (reward)
type Bounty struct {
	Earned   *big.Int
	EndMonth *big.Int
}

// GetRewardEventsInfo retrieves SKALE bounty (reward) events from the ETH blockchain
func GetRewardEventsInfo(fromBlock *big.Int, toBlock *big.Int) {

	var logBountyWasPaidSig = []byte("BountyWasPaid(uint,uint)")
	var logBountyWasPaidSigHash = crypto.Keccak256Hash(logBountyWasPaidSig)
	var TopicsFilter = [][]common.Hash{{logBountyWasPaidSigHash}}

	logs, err := connection.GetBlockLogs(
		fromBlock,
		toBlock,
		TopicsFilter,
		constants.ContractDeploymentAddress[constants.NetActive][constants.Distributor])

	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(binding.DistributorABI)))

	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		var bountyWasPaidEvent BountyWasPaidEvent
		err := contractAbi.Unpack(&bountyWasPaidEvent, "BountyWasPaid", vLog.Data)

		if err != nil {
			log.Fatal(err)
		}
	}
}
