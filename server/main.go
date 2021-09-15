package main

import (
	"fmt"
	"log"
	"net/http"
	"server/router"
	"time"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (s spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("../client/build")).ServeHTTP(w, r)
}

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 4000...")

	spa := spaHandler{staticPath: "../client/build", indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
