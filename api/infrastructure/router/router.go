package router

import (
	"net/http"

	"github.com/Khigashiguchi/khigashiguchi.com/api/interfaces/handlers"
	"github.com/gorilla/mux"
)

// New create http routing.
func New() http.Handler {
	r := mux.NewRouter()

	// create handlers
	getEntries := handlers.NewGetEntriesHandler()

	// create api endpoint
	r.Methods("GET").Path("/api/entries").HandlerFunc(getEntries.Handler)

	// create health check endpoint
	r.Methods("GET").Path("/.health_check").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	return r
}
