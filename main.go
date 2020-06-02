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
	r.HandleFunc("/api/register", api.RegisterHandler).Methods("POST")
	r.HandleFunc("/api/same-calendar", api.SameCalendarHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
	// fmt.Println(samecalendar.SameCalendar(2000, 5))

}
