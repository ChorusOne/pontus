package contract

import (
	"log"
	"math/big"
	"strconv"

	"github.com/ChorusOne/pontus-internal/blockchain"

	"github.com/ChorusOne/pontus-internal/binding"
	"github.com/ChorusOne/pontus-internal/connection"
	"github.com/ChorusOne/pontus-internal/constants"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// PontusValidator wraps ValidatorServiceValidator with a validator id field
type PontusValidator struct {
	ID *big.Int
	binding.ValidatorServiceValidator
	ValidatorEarnings
}

var validatorServiceContractInstance *binding.ValidatorService

// InitValidatorServiceContractInstance returns a SKALE contract instance bound to the active Ethereum network
// TODO: consider design of init methods
func InitValidatorServiceContractInstance() {
	instance, err := binding.NewValidatorService(common.HexToAddress(
		constants.ContractDeploymentAddress[constants.NetActive][constants.ValidatorService]),
		connection.ETHCLIENT)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("SKALE validator service contract is loaded")
	validatorServiceContractInstance = instance
}

// GetValidatorList returns the list of trusted validators for the SKALE network at a given block height
func GetValidatorList(atBlockNumber *big.Int) []*big.Int {
	none := []*big.Int{}
	check := HeightSanityCheck(atBlockNumber, constants.SkaleDeploymentBlockNumber[constants.Mainnet][constants.SkaleToken])

	if !check {
		return none
	}

	validators, err := validatorServiceContractInstance.GetTrustedValidators(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err != nil {
		log.Printf("Could not retrieve validator list: %s\n", err)
		return none
	}
	return validators
}

// GetLatestValidatorList returns the list of trusted validators for the SKALE network at a given block height
func GetLatestValidatorList() []*big.Int {
	none := []*big.Int{}
	lb := blockchain.GetLatestBlockNumber()
	validators, err := validatorServiceContractInstance.GetTrustedValidators(&bind.CallOpts{BlockNumber: lb})

	if err != nil {
		log.Printf("Could not retrieve latest validator list: %s\n", err)
		return none
	}
	return validators
}

// GetValidatorSet returns a set of trusted validators
func GetValidatorSet(validators []*big.Int) map[int64]bool {
	validatorSet := map[int64]bool{}
	for _, validator := range validators {
		validatorSet[validator.Int64()] = true
	}
	return validatorSet
}

// GetLatestValidatorDetails returns details for a given validator ID
func GetLatestValidatorDetails(validator string) binding.ValidatorServiceValidator {

	none := binding.ValidatorServiceValidator{}
	// Current queries at ETH block height
	lb := blockchain.GetLatestBlockNumber()

	// get the list of validators from SKALE
	validatorSet := GetValidatorSet(GetLatestValidatorList())

	//Convert to int
	vi, err := strconv.ParseInt(validator, 10, 64)
	if err != nil {
		log.Printf("Validator id string invalid: %s\n", validator)
		return none
	}

	// check if validator passed is in the set of trusted validators
	_, ok := validatorSet[vi]

	if !ok {
		return none
	}

	v, err := validatorServiceContractInstance.GetValidator(&bind.CallOpts{BlockNumber: lb}, big.NewInt(vi))

	if err != nil {
		log.Printf("Unable to retrieve validator details: %s\n", err)
		return none
	}

	return v
}

// GetLatestValidatorDetailsInt returns details for a given validator ID
// TODO: is returning a single validator useful by itself?
// This is intermediary for embedding in GetAllValidators currently
func GetLatestValidatorDetailsInt(validator *big.Int) binding.ValidatorServiceValidator {

	none := binding.ValidatorServiceValidator{}
	// Current queries at ETH block height
	lb := blockchain.GetLatestBlockNumber()

	// get the list of validators from SKALE
	validatorSet := GetValidatorSet(GetLatestValidatorList())

	// check if validator passed is in the set of trusted validators
	_, ok := validatorSet[validator.Int64()]

	if !ok {
		return none
	}

	v, err := validatorServiceContractInstance.GetValidator(&bind.CallOpts{BlockNumber: lb}, validator)

	if err != nil {
		log.Printf("Unable to retrieve validator details: %s\n", err)
		return none
	}

	return v
}

//GetAllValidatorDetails returns a slice of embedded validator details (with id)
func GetAllValidatorDetails() []PontusValidator {
	pvs := []PontusValidator{}
	vl := GetLatestValidatorList()
	for _, v := range vl {
		pv := PontusValidator{
			ID:                        v,
			ValidatorServiceValidator: GetLatestValidatorDetailsInt(v),
			ValidatorEarnings:         GetValidatorEarnings(v),
		}
		pvs = append(pvs, pv)
	}
	return pvs
}
