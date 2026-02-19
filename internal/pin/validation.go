package pin

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"test/internal/auth"
	"time"
)

func ValidateNegativeShiftLength(ctx context.Context, db *sql.DB, input ClockOut_R) error {
	claims := ctx.Value(auth.ClaimsKey).(*auth.Claims)
	profile_id := claims.ProfileID

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("ValidateNegativeShiftLength: begin tx: %w", err)
	}
	defer tx.Rollback()

	var start_ts time.Time
	err = tx.QueryRowContext(
		ctx,
		`
		SELECT (start_ts)
		FROM shift
		WHERE profile_id = $1
		AND end_ts IS NULL
		`,
		profile_id,
	).Scan(
		&start_ts,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotClockedIn
		}
		return fmt.Errorf("ValidateNegativeShiftLength: db select: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("ValidateNegativeShiftLength: db commit: %w", err)
	}

	if input.EndTs != nil && start_ts.Compare(*input.EndTs) == -1 {
		return ErrNegativeDuration
	}

	return nil
}
