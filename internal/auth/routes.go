package auth

import (
	"database/sql"
	"net/http"
	"test/internal/abstractions"
)

//func RegisterHandler(db *sql.DB) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var input UserCreate
//		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
//			http.Error(w, "invalid body", http.StatusBadRequest)
//			return
//		}
//
//		user, err := CreateUser(r.Context(), db, input)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		json.NewEncoder(w).Encode(user)
//	}
//}
//
//func LoginHandler(db *sql.DB) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var input UserPinAuth
//		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
//			http.Error(w, "invalid body", http.StatusBadRequest)
//			return
//		}
//
//		res, err := ColdStartPin(r.Context(), db, input)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		json.NewEncoder(w).Encode(res)
//	}
//}
//
//func SilentRefreshHandler(db *sql.DB) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var input UserSilentRefresh
//		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
//			http.Error(w, "invalid body", http.StatusBadRequest)
//			return
//		}
//
//		tokens, err := RefreshTokens(r.Context(), db, input)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		json.NewEncoder(w).Encode(tokens)
//	}
//}

//func ReAuthHandler(db *sql.DB) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var input UserReAuth
//		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
//			http.Error(w, "invalid body", http.StatusBadRequest)
//			return
//		}
//
//		tokens, err := WarmStartPin(r.Context(), db, input)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		json.NewEncoder(w).Encode(tokens)
//	}
//}

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
