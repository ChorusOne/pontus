package contract

import (
	"math/big"
)

// InitContracts creates a shared instance of SKALE contract bindings
func InitContracts() {
	InitSkaleTokenContractInstance()
	InitValidatorServiceContractInstance()
	InitDelegationControllerContractInstance()
	InitDistributorContractInstance()
}

// HeightSanityCheck alerts us if we are checking a block before contract deployment
func HeightSanityCheck(blockNumber *big.Int, heightDeployed int64) bool {
	return (blockNumber.Cmp(big.NewInt(heightDeployed)) >= 0)
}
