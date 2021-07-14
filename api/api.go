package api

import (
	"crypto/rand"
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

type BaseHandler struct {
	db *sql.DB
}

func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

type registerResponse struct {
	Key string `json:"key"`
}

func (h *BaseHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	key, err := generateRandomBytes(32)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	keyString := hex.EncodeToString(key)

	err = h.registerApiKey(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(registerResponse{keyString})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

func (h *BaseHandler) UnregisterHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	apiKey, err := hex.DecodeString(q.Get("key"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authenticated, err := h.authenticateApiKey(apiKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if !authenticated {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	err = h.unregisterApiKey(apiKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

type sameCalendarResponse struct {
	Years []int `json:"years"`
}

func (h *BaseHandler) SameCalendarHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	apiKey, err := hex.DecodeString(q.Get("key"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authenticated, err := h.authenticateApiKey(apiKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if !authenticated {
		// Can't authenticate
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Authenticated
	year, err := strconv.Atoi(q.Get("year"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if year < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(q.Get("n"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	forward, err := strconv.ParseBool(q.Get("forward"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	years, err := samecalendar.SameCalendar(year, n, forward)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(sameCalendarResponse{years})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func (h *BaseHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
