package manage

import (
	"database/sql"
	"net/http"
)

func CreateCompanyHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func CreateLocationHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
