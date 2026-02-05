package pin

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func ValidateNegativeShiftLength(ctx context.Context, db *sql.DB, input ClockOut_R) error {
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
		WHERE employment_id = $1
		AND end_ts IS NULL
		`,
		input.EmploymentId,
	).Scan(
		&start_ts,
	)
	if err != nil {
		return fmt.Errorf("ValidateNegativeShiftLength: db select: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("ValidateNegativeShiftLength: db commit: %w", err)
	}

	if input.EndTs.Compare(start_ts) == -1 {
		return ErrNegativeDuration
	}

	return nil
}
