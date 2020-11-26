package live

import (
	"encoding/json"

	"log"
	"net/http"

	"math/big"

	"github.com/ChorusOne/pontus-internal/blockchain"
	"github.com/ChorusOne/pontus-internal/contract"
	"github.com/ChorusOne/pontus-internal/rest/types"
)

// SystemBalanceData requests data for a specific account
func SystemBalanceData() types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		latest := blockchain.GetLatestBlockNumber()
		height := latest
		if heightQuery := r.URL.Query().Get("block_height"); heightQuery != "" {
			height, success := height.SetString(heightQuery, 10)
			// If height is invalid, return details for latest block number
			if !success ||
				height.Cmp(big.NewInt(0)) < 0 ||
				height.Cmp(latest) > 0 {
			}
		}
		w.Header().Set("Content-Type", "application/json")
		accountInfo := getSystemDetails(height)
		if err := json.NewEncoder(w).Encode(accountInfo); err != nil {
			log.Println(err)
		}
	}
}

func getSystemDetails(blockNumber *big.Int) types.System {
	return types.System{
		Block:   blockNumber,
		Balance: contract.SystemSkaleTokenSupply(blockNumber),
	}
}
