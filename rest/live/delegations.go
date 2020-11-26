package live 

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"log"
	"net/http"

	"github.com/ChorusOne/pontus-internal/contract"
	"github.com/ChorusOne/pontus-internal/rest/types"
	"github.com/go-chi/chi"
)

// GetHolderDelegations requests delegations for a holder address
func GetHolderDelegations() types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if account := chi.URLParam(r, "account"); account != "" {
			ha := common.HexToAddress(account)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(contract.GetHolderDelegations(ha)); err != nil {
				log.Printf("Could not encode delegations for holder: %s\n", err)
			}
		}
	}
}

// GetValidatorDelegations requests delegations for a validator id
func GetValidatorDelegations() types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		dtv := big.NewInt(0)
		vs := chi.URLParam(r, "validator")
		vid, ok := dtv.SetString(vs, 10)
		if ok {
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(contract.GetValidatorDelegations(vid)); err != nil {
				log.Printf("Could not encode delegations for holder: %s\n", err)
			}
		}
	}
}
