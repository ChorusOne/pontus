package live 

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"

	"log"
	"net/http"

	"math/big"

	"github.com/ChorusOne/pontus-internal/blockchain"
	"github.com/ChorusOne/pontus-internal/contract"
	"github.com/ChorusOne/pontus-internal/rest/types"
	"github.com/go-chi/chi"
)

// AccountBalanceData requests data for a specific account
func AccountBalanceData() types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if account := chi.URLParam(r, "account"); account != "" {
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
			accountInfo := getAccountDetails(account, height)
			if err := json.NewEncoder(w).Encode(accountInfo); err != nil {
				log.Println(err)
			}
		}
	}
}

func getAccountDetails(address string, blockNumber *big.Int) types.Account {
	return types.Account{
		Address:   common.HexToAddress(address),
		Block:     blockNumber,
		Balance:   contract.AccountSkaleBalance(address, blockNumber),
		Locked:    contract.AccountLockedSkaleBalance(address, blockNumber),
		Delegated: contract.AccountDelegatedSkaleBalance(address, blockNumber)}
}
