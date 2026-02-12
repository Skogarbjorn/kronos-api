package pin

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"test/internal/abstractions"
)

func ClockInHandler(db *sql.DB) http.HandlerFunc {
	return abstractions.CreateJSONHandler(
		db,
		ClockIn,
		WriteDomainError,
	)
}

func ClockOutHandler(db *sql.DB) http.HandlerFunc {
	return abstractions.CreateJSONHandler(
		db,
		ClockOut,
		WriteDomainError,
		ValidateNegativeShiftLength,
	)
}

func ShiftOverviewHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := GetShiftOverview(r.Context(), db)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func ShiftHistoryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := GetShiftHistory(r.Context(), db)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
