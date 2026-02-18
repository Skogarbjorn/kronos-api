package auth

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrProfileNotFound    = errors.New("profile not found")
)

func WriteDomainError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrInvalidCredentials):
		http.Error(w, err.Error(), http.StatusUnauthorized)
	case errors.Is(err, ErrProfileNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
	default:
		log.Printf("internal error: %+v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
