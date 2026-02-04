package manage

import (
	"log"
	"net/http"
)

func WriteDomainError(w http.ResponseWriter, err error) {
	switch {
	default:
		log.Printf("internal error: %+v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
