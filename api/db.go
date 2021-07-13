package api

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./credentials.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func hashKey(key []byte) []byte {
	hash := sha256.New()
	hash.Write(key)
	return hash.Sum(nil)
}

func (h *BaseHandler) CreateCredentialsTable() (err error) {
	_, err = h.db.Exec("CREATE TABLE IF NOT EXISTS credentials (id INTEGER PPRIMARY KEY, apiKey TEXT)")
	if err != nil {
		return err
	}
	return nil
}

func (h *BaseHandler) registerApiKey(key []byte) (err error) {
	hashedKey := hashKey(key)
	hashedKeyString := hex.EncodeToString(hashedKey)

	_, err = h.db.Exec("INSERT INTO credentials (apiKey) VALUES (?)", hashedKeyString)
	if err != nil {
		return err
	}
	return nil
}

func (h *BaseHandler) unregisterApiKey(key []byte) (err error) {
	hashedKey := hashKey(key)
	hashedKeyString := hex.EncodeToString(hashedKey)

	_, err = h.db.Exec("DELETE FROM credentials WHERE apiKey = (?)", hashedKeyString)
	if err != nil {
		return err
	}
	return nil
}

func (h *BaseHandler) authenticateApiKey(key []byte) (authenticated bool, err error) {
	hashedKey := hashKey(key)
	hashedKeyString := hex.EncodeToString(hashedKey)

	row := h.db.QueryRow("SELECT apiKey FROM credentials WHERE apiKey = ?", hashedKeyString)

	var col string
	err = row.Scan(&col)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
