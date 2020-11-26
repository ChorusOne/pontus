package rest

import (
	"database/sql"

	"github.com/ChorusOne/pontus-internal/rest/archive"
	"github.com/ChorusOne/pontus-internal/rest/live"
	"github.com/ChorusOne/pontus-internal/rest/types"

	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/go-chi/chi"
)

var ec *ethclient.Client
var rc *rpc.Client
var db *sql.DB

// StartAPI creates an HTTP server with a REST API for Anthem to consume. This
// function blocks and so should be spawned as a goroutine.
func StartAPI() {
	// Start HTTP router
	r := chi.NewRouter()
	r.Get("/", IndexResponder())
	r.Route("/live", func(r chi.Router) {
		r.Route("/account", func(r chi.Router) {
			r.Route("/{account}", func(r chi.Router) {
				r.Get("/balance", live.AccountBalanceData())
				r.Get(("/delegations"), live.GetHolderDelegations())
			})
		})
		r.Route("/system", func(r chi.Router) {
			r.Get("/", live.SystemBalanceData())
		})
		r.Route("/validator", func(r chi.Router) {
			r.Route("/{validator}", func(r chi.Router) {
				r.Get("/details", live.GetValidatorDetails())
				r.Get("/delegations", live.GetValidatorDelegations())
			})
			r.Get("/list", live.GetCurrentValidatorList())
		})
		r.Route("/validators", func(r chi.Router) {
			r.Get(("/"), live.GetAllValidatorDetails())
		})
	})
	r.Route("/archive", func(r chi.Router) {
		r.Route("/account", func(r chi.Router) {
			r.Route("/{account}", func(r chi.Router) {
				r.Get("/balance", archive.BalanceData())
			})
		})
		r.Route("/system", func(r chi.Router) {
			r.Get("/", archive.SystemBalanceData())
		})
	})
	http.ListenAndServe(":10101", r) // Todo : Config-ify
}

// IndexResponder acts as a health check.
func IndexResponder() types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I am healthy."))
	}
}
