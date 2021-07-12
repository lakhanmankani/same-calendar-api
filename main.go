package main

import (
	"github.com/gorilla/mux"
	"github.com/lakhanmankani/same-calendar-api/api"
	"log"
	"net/http"
)

func main() {
	db := api.ConnectDB()
	h := api.NewBaseHandler(db)

	r := mux.NewRouter()
	r.HandleFunc("/", h.HomeHandler)
	r.HandleFunc("/api/register", h.RegisterHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/same-calendar", h.SameCalendarHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/unregister", h.UnregisterHandler).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", r))
}
