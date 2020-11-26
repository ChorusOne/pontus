package contract

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ChorusOne/pontus-internal/binding"
	"github.com/ChorusOne/pontus-internal/blockchain"
	"github.com/ChorusOne/pontus-internal/connection"
	"github.com/ChorusOne/pontus-internal/constants"
	"github.com/ethereum/go-ethereum/common"
)

// ValidatorEarnings wraps a struct returned from a call to retrieve validator earnings
type ValidatorEarnings struct {
	Earned   *big.Int
	EndMonth *big.Int
}

var distributorContractInstance *binding.Distributor

// InitDistributorServiceContractInstance returns a SKALE contract instance bound to the active Ethereum network
// TODO: consider design of init methods
func InitDistributorContractInstance() {
	instance, err := binding.NewDistributor(common.HexToAddress(
		constants.ContractDeploymentAddress[constants.NetActive][constants.Distributor]),
		connection.ETHCLIENT)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("SKALE distributor service contract is loaded")
	distributorContractInstance = instance
}

// GetValidatorEarnings returns a struct containing the earned fee and end month of said fee.
func GetValidatorEarnings(validator *big.Int) ValidatorEarnings {
	ve, err := distributorContractInstance.GetEarnedFeeAmountOf(&bind.CallOpts{BlockNumber: blockchain.GetLatestBlockNumber()}, validator)
	if err != nil {
		log.Printf("Could not retrieve validator %s earnings: %s\n", validator, err)
	}
	return ValidatorEarnings{ve.Earned, ve.EndMonth}
}
