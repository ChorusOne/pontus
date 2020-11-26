package constants

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Constants for Ethereum and Ethereum test networks
const (
	Rinkeby = "rinkeby"
	Mainnet = "mainnet"
)

// NetActive refers to the SKALE network that Pontus will extract data from
//Defaults to Mainnet.
var NetActive = Mainnet

// GenesisBlockNumber is a mapping from the active Ethereum network to the
// genesis block where SKALE appears.
var GenesisBlockNumber = map[string]*big.Int{
	Rinkeby: big.NewInt(10927184),
	Mainnet: big.NewInt(11060400),
	//first seen delegation accepted event 10927184, 10940028
	//first seen skale transfer 10918200
	//first seen distributor event 10362561
	//first seen delegation event 11100722
}

// Constant names for contracts of interest.
// TODO: Need to clean this up for relevant contracts
const (
	SkaleToken              = "SkaleToken"
	TokenState              = "TokenState"
	DelegationController    = "DelegationController"
	DelegationPeriodManager = "DelegationPeriodManager"
	Bounty                  = "Bounty"
	ValidatorService        = "ValidatorService"
	Distributor             = "Distributor"
	Election                = "Election"
	Governance              = "Governance"
	Validators              = "Validators"
	DowntimeSlasher         = "DowntimeSlasher"
	Accounts                = "Accounts"
	GovernanceSlasher       = "GovernanceSlasher"
)

// ContractDeploymentAddress maps contracts to their deployment addresses
// for different networks of the Skale blockchain
// TODO: Keeo only relevant contracts.
var ContractDeploymentAddress = map[string]map[string]string{

	Rinkeby: {
		SkaleToken:           "0x17fDCB418B9f2f9bddAC98324D5752cE7f707180",
		TokenState:           "0x817EcE46B4A4fF2be3245D80f41d572F970D746F",
		DelegationController: "0x16790939Fd0B4E8c24404Ca3cC5D37C5753d73bd",
		ValidatorService:     "0x1F2157Bf5C820f68826ef1DC71824816Ee795f41",
		Distributor:          "0x2E1102d8b0FD029191aa076a2C1ecA5ab5F6a0AB",
	},

	Mainnet: {
		SkaleToken:           "0x00c83aeCC790e8a4453e5dD3B0B4b3680501a7A7",
		TokenState:           "0x4eE5F270572285776814e32952446e9B7Ee15C86",
		DelegationController: "0x06dD71dAb27C1A3e0B172d53735f00Bf1a66Eb79",
		ValidatorService:     "0x840C8122433A5AA7ad60C1Bcdc36AB9DcCF761a5",
		Distributor:          "0x2a42Ccca55FdE8a9CA2D7f3C66fcddE99B4baB90",
	},
}

// SkaleDeploymentBlockNumber maps contracts to their deployment heights
// for different networks of the SKALE blockchain
// TODO: Update these out for relevant mainnet/rinkeby contracts
var SkaleDeploymentBlockNumber = map[string]map[string]int64{

	Rinkeby: {
		SkaleToken:        6855813,
		Election:          0,
		Governance:        0,
		Validators:        0,
		Accounts:          0,
		GovernanceSlasher: 0,
	},

	Mainnet: {
		SkaleToken:        10363171,
		Election:          0,
		Governance:        0,
		Validators:        0,
		Accounts:          0,
		GovernanceSlasher: 0,
	},
}

//LogTransfer - to capture details from an event log of form Transfer
type LogTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

//LogTransferComment - to capture details from an event log of form TransferComment
type LogTransferComment struct {
	Comment string
}

//LogApproval - to capture details from an event log of form Approval
type LogApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
}

// Constants represting SKALE contract event types
const (
	Minted      = "Minted"
	Transfer    = "Transfer"
	RoleGranted = "RoleGranted"
)

// EventSignature is a map from event type string to byte slice event signature.
var EventSignature = map[string][]byte{
	Minted:      []byte("Minted(address,address,uint256,bytes,bytes)"),
	Transfer:    []byte("Transfer(address,address,uint256)"),
	RoleGranted: []byte("RoleGranted(byte32,address,address"),
}
