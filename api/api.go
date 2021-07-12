package api

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"github.com/lakhanmankani/same-calendar-api/samecalendar"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
	"strconv"
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
	Key string `json:"key"`
}

func registerApiKey(key []byte) {
	h := sha256.New()
	h.Write(key)
	hashedKey := h.Sum(nil)
	hashedKeyString := hex.EncodeToString(hashedKey)

	db, err := sql.Open("sqlite3", "./credentials.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS credentials (id INTEGER PPRIMARY KEY, apiKey TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("INSERT INTO credentials (apiKey) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(hashedKeyString)
	if err != nil {
		log.Fatal(err)
	}
	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func authenticateApiKey(key []byte) (authenticated bool, err error) {
	h := sha256.New()
	h.Write(key)
	hashedKey := h.Sum(nil)
	hashedKeyString := hex.EncodeToString(hashedKey)

	db, err := sql.Open("sqlite3", "./credentials.sqlite")
	if err != nil {
		return false, err
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS credentials (id INTEGER PPRIMARY KEY, apiKey TEXT)")
	if err != nil {
		return false, err
	}
	_, err = stmt.Exec()
	if err != nil {
		return false, err
	}
	err = stmt.Close()
	if err != nil {
		return false, err
	}

	rows, err := db.Query("SELECT apiKey FROM credentials WHERE apiKey = ?", hashedKeyString)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err = rows.Close()
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func unregisterApiKey(key []byte) {
	h := sha256.New()
	h.Write(key)
	hashedKey := h.Sum(nil)
	hashedKeyString := hex.EncodeToString(hashedKey)

	db, err := sql.Open("sqlite3", "./credentials.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS credentials (id INTEGER PPRIMARY KEY, apiKey TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("DELETE FROM credentials WHERE apiKey = (?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(hashedKeyString)
	if err != nil {
		log.Fatal(err)
	}
}


func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	key, err := generateRandomBytes(32)

	h := sha256.New()
	h.Write(key)
	apiKey := hex.EncodeToString(h.Sum(nil))
	registerApiKey(h.Sum(nil))

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(Register{apiKey})
	if err != nil {
		log.Fatal(err)
	}
}

func UnregisterHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	apiKey, err := hex.DecodeString(q.Get("key"))
	if err != nil {
		log.Fatal(err)
	}
	authenticated, err := authenticateApiKey(apiKey)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Fatal(err)
	}
	if !authenticated {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	unregisterApiKey(apiKey)
}

type SameCalendar struct {
	Years []string `json:"years"`
}

func SameCalendarHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	apiKey, err := hex.DecodeString(q.Get("key"))
	if err != nil {
		log.Fatal(err)
	}
	authenticated, err := authenticateApiKey(apiKey)
	if !authenticated {
		// Can't authenticate
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Authenticated
	year, err := strconv.Atoi(q.Get("year"))
	if err != nil {
		log.Fatal(err)
	}

	n, err := strconv.Atoi(q.Get("n"))
	if err != nil {
		log.Fatal(err)
	}
	years, err := samecalendar.SameCalendar(year, n)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(years)
	if err != nil {
		log.Fatal(err)
	}
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
