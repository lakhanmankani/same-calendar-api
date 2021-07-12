package main

import (
	"github.com/gorilla/mux"
	"github.com/lakhanmankani/same-calendar-api/api"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", api.HomeHandler)
	r.HandleFunc("/api/register", api.RegisterHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/same-calendar", api.SameCalendarHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/unregister", api.UnregisterHandler).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", r))
}
