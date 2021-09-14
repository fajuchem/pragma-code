package main

import (
	"fmt"
	"log"
	"net/http"
	"server/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 4000...")

	log.Fatal(http.ListenAndServe(":4000", r))
}
