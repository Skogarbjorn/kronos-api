package pin

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
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

func SyncShiftHandler(db *sql.DB) http.HandlerFunc {
	return abstractions.CreateJSONHandler(
		db,
		SyncShift,
		WriteDomainError,
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

func GetLocationsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := GetLocations(r.Context(), db)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func GetTasksHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s_location_id := r.URL.Query().Get("location_id")
		var location_id *int

		if s_location_id != "" {
			parsed, err := strconv.Atoi(s_location_id)
			if err != nil {
				http.Error(w, "location_id must be an integer", http.StatusBadRequest)
				return
			}
			location_id = &parsed
		}

		result, err := GetTasks(r.Context(), db, location_id)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func GetEmploymentsDetailedHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := GetEmploymentsDetailed(r.Context(), db)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
