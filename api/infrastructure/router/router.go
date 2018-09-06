package router

import (
	"net/http"

	"github.com/Khigashiguchi/khigashiguchi.com/api/infrastructure/repository"
	"github.com/Khigashiguchi/khigashiguchi.com/api/interfaces/handlers"
	"github.com/gorilla/mux"
)

// New create http routing.
func New(db repository.DBConnector) http.Handler {
	r := mux.NewRouter()

	// create api endpoint
	r.Methods("GET").Path("/api/entries").HandlerFunc(handlers.NewGetEntriesHandler(db).Handler)
	r.Methods("POST").Path("/api/entries").HandlerFunc(handlers.NewPostEntriesHandler(db).Handler)

	// create health check endpoint
	r.Methods("GET").Path("/.health_check").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	return r
}
