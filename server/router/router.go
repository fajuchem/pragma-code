package router

import (
	"server/beer"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/products", beer.GetAllProducts).Methods("GET", "OPTIONS")
	return router
}
