package auth

import (
	"database/sql"
	"net/http"
	"test/internal/abstractions"
)

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return abstractions.CreateJSONHandler(db, CreateUser, WriteDomainError)
}

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return abstractions.CreateJSONHandler(db, ColdStartPin, WriteDomainError)
}

func SilentRefreshHandler(db *sql.DB) http.HandlerFunc {
	return abstractions.CreateJSONHandler(db, RefreshTokens, WriteDomainError)
}

func ReAuthHandler(db *sql.DB) http.HandlerFunc {
	return abstractions.CreateJSONHandler(db, WarmStartPin, WriteDomainError)
}
