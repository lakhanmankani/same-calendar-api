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
	stmt, err := h.db.Prepare("CREATE TABLE IF NOT EXISTS credentials (id INTEGER PPRIMARY KEY, apiKey TEXT)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (h *BaseHandler) registerApiKey(key []byte) (err error) {
	hashedKey := hashKey(key)
	hashedKeyString := hex.EncodeToString(hashedKey)

	stmt, err := h.db.Prepare("INSERT INTO credentials (apiKey) VALUES (?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(hashedKeyString)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (h *BaseHandler) unregisterApiKey(key []byte) (err error) {
	hashedKey := hashKey(key)
	hashedKeyString := hex.EncodeToString(hashedKey)

	stmt, err := h.db.Prepare("DELETE FROM credentials WHERE apiKey = (?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(hashedKeyString)
	if err != nil {
		return err
	}
	return nil
}

func (h *BaseHandler) authenticateApiKey(key []byte) (authenticated bool, err error) {
	hashedKey := hashKey(key)
	hashedKeyString := hex.EncodeToString(hashedKey)

	rows, err := h.db.Query("SELECT apiKey FROM credentials WHERE apiKey = ?", hashedKeyString)
	if err != nil {
		return false, err
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
