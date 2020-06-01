package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

type Register struct {
	ApiKey string `json:"api_key"`
}

func apiRegisterHandler(w http.ResponseWriter, r *http.Request) {
	key, err := GenerateRandomBytes(32)

	h := sha256.New()
	h.Write(key)
	apiKey := hex.EncodeToString(h.Sum(nil))
	fmt.Println(apiKey)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(Register{apiKey})
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Write api key to database
}
func apiSameCalendarHandler(w http.ResponseWriter, r *http.Request) {

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	contents := []byte("<h1>Home</h1>")
	_, _ = w.Write(contents)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/api/register", apiRegisterHandler).Methods("POST")
	r.HandleFunc("/api/samecalendar", apiSameCalendarHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
	// fmt.Println(samecalendar.SameCalendar(2000, 5))

}
