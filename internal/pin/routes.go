package pin

import (
	"database/sql"
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
