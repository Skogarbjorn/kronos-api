package abstractions

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
)

type CreatorFunc[I any, O any] func(
	ctx context.Context,
	db *sql.DB,
	input I,
) (O, error)

func CreateJSONHandler[I any, O any](
	db *sql.DB,
	create CreatorFunc[I, O],
	writeError ErrorWriter,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input I

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}

		result, err := create(r.Context(), db, input)
		if err != nil {
			writeError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
