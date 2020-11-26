package live

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/ChorusOne/pontus-internal/contract"
	"github.com/ChorusOne/pontus-internal/rest/types"
	"github.com/go-chi/chi"
)

// GetCurrentValidatorList returns the list of current validators
func GetCurrentValidatorList() types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vl := contract.GetLatestValidatorList()
		if err := json.NewEncoder(w).Encode(vl); err != nil {
			log.Printf("Could not encode validator list: %s\n", err)
		}
	}
}

// GetValidatorDetails requests data for a specific validator ID
func GetValidatorDetails() types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if validator := chi.URLParam(r, "validator"); validator != "" {
			vd := contract.GetLatestValidatorDetails(validator)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(vd); err != nil {
				log.Printf("Could not encode validator details: %s\n", err)
			}

		}
	}
}

// GetAllValidatorDetails requests data for a specific validator ID
func GetAllValidatorDetails() types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		vds := contract.GetAllValidatorDetails()
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(vds); err != nil {
			log.Printf("Could not encode validator details: %s\n", err)
		}

	}
}
