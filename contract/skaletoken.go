package contract

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ChorusOne/pontus-internal/binding"
	"github.com/ChorusOne/pontus-internal/connection"
	"github.com/ChorusOne/pontus-internal/constants"
)

var skaleTokenContractInstance *binding.SkaleToken

// InitSkaleTokenContractInstance returns a SKALE contract instance bound to the active Ethereum network
// TODO: consider design of init methods, cleanup
func InitSkaleTokenContractInstance() {

	instance, err := binding.NewSkaleToken(common.HexToAddress(
		constants.ContractDeploymentAddress[constants.NetActive][constants.SkaleToken]),
		connection.ETHCLIENT)

	if err != nil {
		log.Fatalf("Could not load SKALE token contract instance: %s\n", err)
	}
	log.Println("SKALE token contract is loaded")
	skaleTokenContractInstance = instance
}

//AccountSkaleBalance returns the SKALE balance for a given address at a particular block height
func AccountSkaleBalance(accountAddressHex string, atBlockNumber *big.Int) *big.Int {

	check := HeightSanityCheck(atBlockNumber,
		constants.SkaleDeploymentBlockNumber[constants.Mainnet][constants.SkaleToken])

	if !check {
		zero := big.NewInt(0)
		return zero
	}

	holder := common.HexToAddress(accountAddressHex)
	g, err := skaleTokenContractInstance.BalanceOf(&bind.CallOpts{BlockNumber: atBlockNumber}, holder)

	if err != nil {
		log.Printf("Could not read SKALE token balance at %s: %s\n", atBlockNumber, err)
	}
	return g
}

// SystemSkaleTokenSupply returns the total number of SKALE tokens at a given block number
func SystemSkaleTokenSupply(atBlockNumber *big.Int) *big.Int {

	ts := big.NewInt(0)

	check := HeightSanityCheck(atBlockNumber,
		constants.SkaleDeploymentBlockNumber[constants.NetActive][constants.SkaleToken])

	if !check {
		return ts
	}

	ts, err := skaleTokenContractInstance.TotalSupply(&bind.CallOpts{BlockNumber: atBlockNumber})

	if err != nil {
		log.Printf("Could not read SKALE system supply: %s at block number: %s\n", err, atBlockNumber)
	}
	return ts
}

// AccountLockedSkaleBalance returns the total amount of locked SKALE tokens for an account at a given block height
func AccountLockedSkaleBalance(accountAddressHex string, atBlockNumber *big.Int) *big.Int {
	zero := big.NewInt(0)
	check := HeightSanityCheck(atBlockNumber, constants.SkaleDeploymentBlockNumber[constants.Mainnet][constants.SkaleToken])

	if !check {
		return zero
	}

	account := common.HexToAddress(accountAddressHex)
	ls, err := skaleTokenContractInstance.GetAndUpdateLockedAmountReadOnly(&bind.CallOpts{BlockNumber: atBlockNumber}, account)

	if err != nil {
		log.Printf("Error reading SKALE locked balance: %s, address: %s, block number: %s \n", err, accountAddressHex, atBlockNumber)
		return zero
	}

	return ls
}

// AccountDelegatedSkaleBalance returns the supply of delegated tokens for a given account
func AccountDelegatedSkaleBalance(accountAddressHex string, atBlockNumber *big.Int) *big.Int {
	zero := big.NewInt(0)
	check := HeightSanityCheck(atBlockNumber, constants.SkaleDeploymentBlockNumber[constants.Mainnet][constants.SkaleToken])

	if !check {
		return zero
	}

	account := common.HexToAddress(accountAddressHex)
	ds, err := skaleTokenContractInstance.GetAndUpdateDelegatedAmountReadOnly(&bind.CallOpts{BlockNumber: atBlockNumber}, account)

	if err != nil {
		log.Printf("Error reading SKALE delegated balance: %s, address: %s, block number: %s \n", err, accountAddressHex, atBlockNumber)
	}

	return ds
}

// AccountSlashedSkaleBalance returns the supply of delegated tokens for a given account
func AccountSlashedSkaleBalance(accountAddressHex string, atBlockNumber *big.Int) *big.Int {
	zero := big.NewInt(0)
	check := HeightSanityCheck(atBlockNumber, constants.SkaleDeploymentBlockNumber[constants.Mainnet][constants.SkaleToken])

	if !check {
		return zero
	}

	account := common.HexToAddress(accountAddressHex)
	ss, err := skaleTokenContractInstance.GetAndUpdateSlashedAmountReadOnly(&bind.CallOpts{BlockNumber: atBlockNumber}, account)

	if err != nil {
		log.Printf("Error reading SKALE slashed balance: %s, address: %s, block number: %s \n", err, accountAddressHex, atBlockNumber)
		return zero
	}

	return ss
}
