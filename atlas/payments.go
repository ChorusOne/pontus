package atlas

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Atlas Functions relevant to validator and group epoch payments (aka commissions, in cUSD)

//EpochPaymentEvent is a struct to represent data for the event
//ValidatorEpochPaymentDistributed(address,unit256,address,uint256)
//emitted by Validators.sol
type EpochPaymentEvent struct {
	Validator        common.Address
	ValidatorPayment *big.Int
	Group            common.Address
	GroupPayment     *big.Int
	BlockNumber      *big.Int
}

//GetEpochPaymentEvents fetches data for all
//ValidatorEpochPaymentDistributed events emitted by the Validators contract
// between (and including) the two specified block numbers.
/* func GetEpochPaymentEvents(fromBlock *big.Int, toBlock *big.Int) []*EpochPaymentEvent {

	var logEpochPaymentSig = []byte("ValidatorEpochPaymentDistributed(address,uint256,address,uint256)")
	var logEpochPaymentSigHash = crypto.Keccak256Hash(logEpochPaymentSig)
	var TopicsFilter = [][]common.Hash{{logEpochPaymentSigHash}}

	contractAddress := common.HexToAddress(WrapperContractDeploymentAddress[NetActive][Validators])
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Topics:    TopicsFilter,

		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := atlasEthClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	paymentsDistributed := make([]*EpochPaymentEvent, 0, len(logs))

	contractAbi, err := abi.JSON(strings.NewReader(string(binding.ValidatorsABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {

		var epochPayment EpochPaymentEvent
		err := contractAbi.Unpack(&epochPayment, "ValidatorEpochPaymentDistributed", vLog.Data)

		if err != nil {
			log.Fatal(err)
		}

		epochPayment.Validator = common.HexToAddress(vLog.Topics[1].Hex())
		epochPayment.Group = common.HexToAddress(vLog.Topics[2].Hex())
		epochPayment.BlockNumber = new(big.Int).SetUint64(vLog.BlockNumber)

		paymentsDistributed = append(paymentsDistributed, &epochPayment)
	}

	return paymentsDistributed
} */
