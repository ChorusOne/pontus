package contract

import (
	"log"
	"math/big"

	"github.com/ChorusOne/pontus-internal/blockchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ChorusOne/pontus-internal/binding"
	"github.com/ChorusOne/pontus-internal/connection"
	"github.com/ChorusOne/pontus-internal/constants"
)

var delegationContractInstance *binding.DelegationController

// InitDelegationControllerContractInstance returns a SKALE contract instance bound to the active Ethereum network
// TODO: consider design of init methods, cleanup
func InitDelegationControllerContractInstance() {

	instance, err := binding.NewDelegationController(common.HexToAddress(
		constants.ContractDeploymentAddress[constants.NetActive][constants.DelegationController]),
		connection.ETHCLIENT)

	if err != nil {
		log.Fatalf("Could not load SKALE delegation controller contract instance: %s\n", err)
	}
	log.Println("SKALE delegation controller contract is loaded")
	delegationContractInstance = instance
}

// GetHolderDelegations returns the SKALE delegation information for a SKALE holder address
func GetHolderDelegations(holder common.Address) []binding.DelegationControllerDelegation {
	lb := blockchain.GetLatestBlockNumber()
	nd, err := delegationContractInstance.GetDelegationsByHolderLength(&bind.CallOpts{BlockNumber: lb}, holder)

	if err != nil {
		log.Printf("Error retrieving holder elegations length: %s\n", err)
	}

	ndi := nd.Int64()
	hdiList := []binding.DelegationControllerDelegation{}

	var d int64
	for d = 0; d < ndi; d++ {
		dbi := big.NewInt(d)
		vd, err := delegationContractInstance.DelegationsByHolder(&bind.CallOpts{BlockNumber: lb}, holder, dbi)
		if err != nil {
			log.Printf("Error retriving validator %s delegation number: %s : %s\n", holder, vd, err)
		}
		vdi, err := delegationContractInstance.GetDelegation(&bind.CallOpts{BlockNumber: lb}, vd)
		if err != nil {
			log.Printf("Error retriving validator %s: delegation number: %s info %s\n", holder, vd, err)
		}
		hdiList = append(hdiList, vdi)
	}
	return hdiList
}

// GetValidatorDelegations returns the SKALE delegation information for a SKALE validator ID
func GetValidatorDelegations(validatorID *big.Int) []binding.DelegationControllerDelegation {
	lb := blockchain.GetLatestBlockNumber()
	nd, err := delegationContractInstance.GetDelegationsByValidatorLength(&bind.CallOpts{BlockNumber: lb}, validatorID)

	if err != nil {
		log.Printf("Error retrieving validator delegations length: %s\n", err)
	}

	ndi := nd.Int64()
	vdiList := []binding.DelegationControllerDelegation{}

	var d int64
	for d = 0; d < ndi; d++ {
		dbi := big.NewInt(d)
		vd, err := delegationContractInstance.DelegationsByValidator(&bind.CallOpts{BlockNumber: lb}, validatorID, dbi)
		if err != nil {
			log.Printf("Error retriving validator %s delegation number: %s : %s\n", validatorID, vd, err)
		}
		vdi, err := delegationContractInstance.GetDelegation(&bind.CallOpts{BlockNumber: lb}, vd)
		if err != nil {
			log.Printf("Error retriving validator %s: delegation number: %s info %s\n", validatorID, vd, err)
		}
		vdiList = append(vdiList, vdi)
	}
	return vdiList
}

/*
func GetDelegationsToValidatorNow(holder common.Address, validator *big.Int) *big.Int {
	lb := blockchain.GetLatestBlockNumber()
	dtv, err := delegationContractInstance.GetAndUpdateDelegatedByHolderToValidatorNowReadOnly(&bind.CallOpts{BlockNumber: lb}, holder, validator)
	if err != nil {
		log.Printf("Could not retrieve validator delegations from holder %s: to validator: %s.  %s\n", validator, holder, err)
	}
	log.Printf("Delegations from holder: %s to validator: %s are: %s", holder, validator, dtv)
	return dtv
} */
