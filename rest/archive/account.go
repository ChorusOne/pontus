package archive

import (
	"encoding/json"
	"log"

	"github.com/ChorusOne/pontus-internal/rest/types"
	"github.com/ChorusOne/pontus-internal/snapshot"
	"github.com/go-chi/chi"

	"net/http"
)

// BalanceData returns data for a specific account at all snapshot points
func BalanceData() types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		// Handle accountID = ""
		if account := chi.URLParam(r, "account"); account != "" {

			datedAccounts, _ := getDetails(account)

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(datedAccounts); err != nil {
				log.Println(err)

			}

		}
	}
}

func getDetails(account string) ([]*types.DatedAccount, error) {
	snapshotDates, err := snapshot.DatesOfCompletedSnapshots()

	if err != nil {
		return nil, err
	}

	var datedAccounts = make([]*types.DatedAccount, 0, len(snapshotDates))

	// Fetch account rows from snapshot_data table
	accountSnapshot, err := snapshot.FetchAccountSnapshot(account)
	if err != nil {
		return nil, err
	}

	for _, snapshotDate := range snapshotDates {
		datedAccount := types.NewDatedAccount(snapshotDate.Date, account)
		datedAccount.Height = snapshotDate.BlockNumber

		if asr, ok := accountSnapshot[snapshotDate.BlockNumber]; ok {
			datedAccount.SkaleTokenBalance = asr.SkaleTokenBalance
			datedAccount.SkaleTokenLockedBalance = asr.SkaleTokenLockedBalance
			datedAccount.SkaleTokenUSDValue = "0"
		}
		// Fetch Rewards Data and Add To Account
		datedAccounts = append(datedAccounts, datedAccount)
	}
	return datedAccounts, nil
}
