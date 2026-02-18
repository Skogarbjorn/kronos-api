package pin

import (
	"errors"
	"log"
	"net/http"

	"github.com/lib/pq"
)

var (
	ErrShiftAlreadyExists = errors.New("already clocked in")
	ErrNotClockedIn       = errors.New("not clocked in")
	ErrNegativeDuration   = errors.New("shift duration cannot be negative")
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
	case errors.Is(err, ErrNegativeDuration):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		log.Printf("internal error: %+v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
