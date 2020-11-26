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

// DelegationAcceptedEvent stores information representing data for the event DelegationAccepted(uint)
type DelegationAcceptedEvent struct {
	DelegationID *big.Int
}

// Delegation represents SKALE delegation data
type Delegation struct {
	Holder           common.Address // address of token owner
	ValidatorID      *big.Int
	Amount           *big.Int
	DelegationPeriod *big.Int
	Created          *big.Int // time of delegation creation
	Started          *big.Int // month when a delegation becomes active
	Finished         *big.Int // first month after a delegation ends
	Info             string
}

// GetDelegationEventsInfo retrieves SKALE delegation events from the ETH blockchain
func GetDelegationEventsInfo(fromBlock *big.Int, toBlock *big.Int) {

	var logDelegationAcceptedSig = []byte("DelegationAccepted(uint256)")
	var logDelegationProposedSig = []byte("DelegationProposed(uint256)")
	var logDelegationRequestCanceledByUserSig = []byte("DelegationRequestCanceledByUser(uint256)")
	var logUndelegationRequestedSig = []byte("UndelegationRequested(uint256)")

	var logDelegationAcceptedSigHash = crypto.Keccak256Hash(logDelegationAcceptedSig)
	var logDelegationProposedSigHash = crypto.Keccak256Hash(logDelegationProposedSig)
	var logDelegationRequestCanceledByUserSigHash = crypto.Keccak256Hash(logDelegationRequestCanceledByUserSig)
	var logUndelegationRequestedSigHash = crypto.Keccak256Hash(logUndelegationRequestedSig)

	log.Println(logDelegationAcceptedSigHash.Hex())
	log.Println(logDelegationProposedSigHash.Hex())
	log.Println(logDelegationRequestCanceledByUserSigHash.Hex())
	log.Println(logUndelegationRequestedSigHash.Hex())

	var TopicsFilter = [][]common.Hash{{logDelegationAcceptedSigHash}}

	logs, err := connection.GetBlockLogs(
		fromBlock,
		toBlock,
		TopicsFilter,
		constants.ContractDeploymentAddress[constants.NetActive][constants.DelegationController])

	if err != nil {
		log.Fatal(err)
	}

	//rewardsInfo := make([]*RewardInfo, 0, len(logs))
	contractAbi, err := abi.JSON(strings.NewReader(string(binding.DelegationControllerABI)))

	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		var delegationAcceptedEvent DelegationAcceptedEvent
		err := contractAbi.Unpack(&delegationAcceptedEvent, "DelegationAccepted", vLog.Data)

		if err != nil {
			log.Fatal(err)
		}
	}
}
