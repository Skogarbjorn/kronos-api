CREATE UNIQUE INDEX one_ongoing_shift_per_employment
ON shift (profile_id)
WHERE end_ts IS NULL;
