package types

import (
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

// Useful Aliases
// -----------------------------------------------------------------------------

// Address is an alias of string
type Address = string

// Height of the chain is a simple integer.
type Height = string

// Amount is any balance or vote - an abstraction over SKALE quantification
type Amount = string

// Date is an alias for string (for now)
type Date = string

// Name is an alias for a validator name
type Name string

// Description is an alias for a validator description
type Description string

// Message is an alias for a response string
type Message = string

// Concrete API Types
// ----------------------------------------------------------------------------

// Account is a type that contains all the relevant information that Pontus seeks for an account
type Account struct {
	Address   common.Address
	Block     *big.Int
	Balance   *big.Int
	Locked    *big.Int
	Delegated *big.Int
}

// System is a type that contains all the relevant information that Pontus seeks for the system
type System struct {
	Block   *big.Int
	Balance *big.Int
}

//DatedAccount represents an account and its details extracted from a daily snapshot
type DatedAccount struct {
	SnapshotDate       Date    `json:"snapshotDate"`
	Address            Address `json:"address"`
	Height             Height  `json:"height"`
	SnapshotReward     Amount  `json:"snapshotReward"`
	SnapshotCommission Amount  `json:"snapshotCommission"`

	SkaleTokenBalance       Amount `json:"skaleTokenBalance"`
	SkaleTokenLockedBalance Amount `json:"skaleTokenLockedBalance"`
	SkaleTokenUSDValue      Amount `json:"skaleTokenUSDValue"`
	//NonVotingLockedSkaleBalance Amount `json:"nonVotingLockedGoldBalance"`
	//VotingLockedSkaleBalance    Amount `json:"votingLockedGoldBalance"`
	//PendingWithdrawalBalance    Amount `json:"pendingWithdrawalBalance"`
	//Delegations []*Delegation `json:"delegations"`
}

//NewDatedAccount is a constructor for DatedAccount
func NewDatedAccount(date string, address string) *DatedAccount {

	return &DatedAccount{
		SnapshotDate:       date,
		Address:            address,
		Height:             "0",
		SnapshotReward:     "0",
		SnapshotCommission: "0",

		SkaleTokenBalance:       "0",
		SkaleTokenLockedBalance: "0",
		SkaleTokenUSDValue:      "0",
		// NonVotingLockedSkaleBalance: "0",
		// VotingLockedSkaleBalance: "0",
		// PendingWithdrawalBalance: "0",
		// Delegations: []*Delegation{},
	}

}

// InvalidRequest provides JSON details for an invalid request
type InvalidRequest struct {
	Message Message `json:"message"`
}

// NewInvalidRequest is a constructor for an invalid request
func NewInvalidRequest(message string) *InvalidRequest {
	return &InvalidRequest{
		Message: message,
	}
}

// Handler is a small type that helps to make returning closures in the functions
// below less verbose.
type Handler = func(w http.ResponseWriter, r *http.Request)
