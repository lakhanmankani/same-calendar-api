package api

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func generateRandomBytes(n int) ([]byte, error) {
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

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	key, err := generateRandomBytes(32)

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

func SameCalendarHandler(w http.ResponseWriter, r *http.Request) {

}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	contents := []byte("<h1>Home</h1>")
	_, _ = w.Write(contents)
}
