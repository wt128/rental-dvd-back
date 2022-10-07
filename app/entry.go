package main

import (
	"fmt"
	"goapi/controller"
	"net/http"
	"strings"
)

func RentalEntries(w http.ResponseWriter, r *http.Request) {
	/* TODO: 複数のパスに対応させる*/
	path := strings.Split(r.URL.Path, "/")[2]
	switch r.Method {
	case http.MethodGet:
		switch path {
		case "":
			controller.RentalAll(w, r)
		}
	case http.MethodPost:

	}

}

func ActorEntries(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			controller.RentalAll(w, r)
		}
	}
}
