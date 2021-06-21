package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/isuryanarayanan/go-rest-api/router"
	"github.com/isuryanarayanan/go-rest-api/settings"
)

func main() {
	settings.ConnectDatabase()

	r := mux.NewRouter()

	// Book routing
	router.BookRouter(r)
	// User routing
	router.UserRouter(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
