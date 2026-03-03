package db

import (
	"database/sql"
	"log"
	"os"
)

func create(db *sql.DB, schema string) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	log.Println("Table created successfully.")
}
func insert(db *sql.DB, row string) {
	_, err := db.Exec(row)
	if err != nil {
		log.Fatalf("Failed to insert row: %v", err)
	}
	log.Println("Row inserted successfully.")
}

func CreateTables(db *sql.DB) {
	profile := `
	CREATE TABLE IF NOT EXISTS profile (
		id SERIAL PRIMARY KEY,
		kt VARCHAR(10) UNIQUE NOT NULL,
		first_name VARCHAR(128),
		last_name VARCHAR(128),
		created TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`
	profile_pin_auth := `
	CREATE TABLE IF NOT EXISTS profile_pin_auth (
		profile_id INT NOT NULL,
		pin TEXT NOT NULL,
		created TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (profile_id) REFERENCES profile(id)
	);`
	profile_password_auth := `
	CREATE TABLE IF NOT EXISTS profile_password_auth (
		profile_id INT NOT NULL,
		email VARCHAR(128) NOT NULL,
		password TEXT NOT NULL,
		created TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (profile_id) REFERENCES profile(id)
	);`
	workspace := `
	CREATE TABLE IF NOT EXISTS workspace (
		id SERIAL PRIMARY KEY,
		name VARCHAR(128) NOT NULL,
		created TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`
	location := `
	CREATE TABLE IF NOT EXISTS location (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		address TEXT NOT NULL,
		workspace_id INT NOT NULL,
		created TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (workspace_id) REFERENCES workspace(id)
	);`
	company := `
	CREATE TABLE IF NOT EXISTS company (
		id SERIAL PRIMARY KEY,
		name VARCHAR(128) NOT NULL,
		workspace_id INT NOT NULL,
		created TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (workspace_id) REFERENCES workspace(id)
	);`
	contract := `
	CREATE TABLE IF NOT EXISTS contract (
		id SERIAL PRIMARY KEY,
	    hourly_rate INT,
	    unpaid_lunch_minutes INT,
		created TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`
	employment := `
	CREATE TABLE IF NOT EXISTS employment (
		id SERIAL PRIMARY KEY,
		profile_id INT,
		company_id INT,
		contract_id INT,
		role VARCHAR(20) CHECK (role IN ('admin', 'manager', 'worker')),
		start_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		end_date TIMESTAMPTZ,
		created TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (profile_id) REFERENCES profile(id),
		FOREIGN KEY (company_id) REFERENCES company(id),
		FOREIGN KEY (contract_id) REFERENCES contract(id)
	);`
	task := `
	CREATE TABLE IF NOT EXISTS task (
		id SERIAL PRIMARY KEY,
		location_id INT NOT NULL,
		company_id INT NOT NULL,
		name TEXT NOT NULL,
		description TEXT,
		is_completed BOOLEAN DEFAULT FALSE,
		created TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (company_id) REFERENCES company(id) ON DELETE CASCADE,
		FOREIGN KEY (location_id) REFERENCES location(id) ON DELETE CASCADE
	);`
	shift := `
	CREATE TABLE IF NOT EXISTS shift (
		id SERIAL PRIMARY KEY,
		profile_id INT NOT NULL,
		task_id INT NOT NULL,
		start_ts TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		end_ts TIMESTAMPTZ,
		created TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (profile_id) REFERENCES profile(id),
		FOREIGN KEY (task_id) REFERENCES task(id)
	);`
	refresh_token := `
	CREATE TABLE IF NOT EXISTS refresh_token (
		id SERIAL PRIMARY KEY,
		profile_id INT NOT NULL,
		device_id TEXT NOT NULL,
		token_hash TEXT NOT NULL,
		expires_at TIMESTAMP NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE,
		UNIQUE(profile_id, device_id)
	);`

	create(db, profile)
	create(db, profile_pin_auth)
	create(db, profile_password_auth)
	create(db, workspace)
	create(db, location)
	create(db, company)
	create(db, contract)
	create(db, employment)
	create(db, task)
	create(db, shift)
	create(db, refresh_token)
}

func DropTables(db *sql.DB) {
	drop_sql := `
	DROP SCHEMA public CASCADE;
	CREATE SCHEMA public;
	GRANT ALL ON SCHEMA public TO postgres;
	GRANT ALL ON SCHEMA public TO public;`

	_, err := db.Exec(drop_sql)
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}
	log.Println("Tables dropped successfully.")
}

func InsertDummy(db *sql.DB) {
	b, err := os.ReadFile("internal/db/insert.sql")
	_, err = db.Exec(string(b))
	if err != nil {
		log.Fatalf("Failed insert db actions: %v", err)
	}
	log.Println("Insert db actions successful")
}

func MiscDB(db *sql.DB) {
	b, err := os.ReadFile("internal/db/misc.sql")
	_, err = db.Exec(string(b))
	if err != nil {
		log.Fatalf("Failed misc db actions: %v", err)
	}
	log.Println("Miscellaneous db actions successful")
}
