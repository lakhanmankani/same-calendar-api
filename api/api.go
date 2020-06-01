package api

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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

func registerApiKey(key []byte) {
	h := sha256.New()
	h.Write(key)
	hashedKey := h.Sum(nil)
	hashedKeyString := hex.EncodeToString(hashedKey)

	db, err:= sql.Open("sqlite3", "./credentials.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS keys (id INTEGER PPRIMARY KEY, apiKey TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("INSERT INTO keys (apiKey) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(hashedKeyString)
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	key, err := generateRandomBytes(32)

	h := sha256.New()
	h.Write(key)
	apiKey := hex.EncodeToString(h.Sum(nil))
	fmt.Println(apiKey)
	registerApiKey(h.Sum(nil))

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
