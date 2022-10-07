package main

import (
	"net/http"
)

func Router(mux *http.ServeMux) {
	mux.HandleFunc("/rental/", RentalEntries)
	mux.HandleFunc("/actor/", ActorEntries)

}
