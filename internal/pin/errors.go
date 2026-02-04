package pin

import (
	"errors"
	"log"
	"net/http"

	"github.com/lib/pq"
)

var (
	ErrShiftAlreadyExists = errors.New("user already clocked in")
	ErrNotClockedIn       = errors.New("user is not clocked in")
)

func translateDBError(err error) error {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		case "23505":
			if pqErr.Constraint == "one_ongoing_shift_per_employment" {
				return ErrShiftAlreadyExists
			}
		}
	}
	return err
}

func WriteDomainError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrShiftAlreadyExists):
		http.Error(w, err.Error(), http.StatusConflict)
	case errors.Is(err, ErrNotClockedIn):
		http.Error(w, err.Error(), http.StatusConflict)
	default:
		log.Printf("internal error: %+v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
