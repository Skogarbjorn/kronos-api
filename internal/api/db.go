package api

import (
	"database/sql"
	"log"
)

func create(db *sql.DB, schema string) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	log.Println("Table created successfully.")
}

func CreateTables(db *sql.DB) {
	create_user := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		kt VARCHAR(10) UNIQUE,
		first_name VARCHAR(128),
		last_name VARCHAR(128)
	);`
	create_user_pin_auth := `
	CREATE TABLE IF NOT EXISTS user_pin_auth (
		user_id INT,
		pin TEXT,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`
	create_user_password_auth := `
	CREATE TABLE IF NOT EXISTS user_password_auth (
		user_id INT,
		email VARCHAR(128),
		password TEXT,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`
	create_company := `
	CREATE TABLE IF NOT EXISTS company (
		id SERIAL PRIMARY KEY,
		name VARCHAR(128)
	);`
	create_contract := `
	CREATE TABLE IF NOT EXISTS contract (
		id SERIAL PRIMARY KEY,
		type VARCHAR(20) CHECK (type IN ('employee', 'contracter')),
	    hourly_rate INT,
	    unpaid_lunch_minutes INT
	);`
	create_employment := `
	CREATE TABLE IF NOT EXISTS employment (
		id SERIAL PRIMARY KEY,
		user_id INT,
		company_id INT,
		contract_id INT,
		role VARCHAR(20) CHECK (role IN ('admin', 'manager', 'worker')),
		start_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		end_date TIMESTAMPTZ,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (company_id) REFERENCES company(id),
		FOREIGN KEY (contract_id) REFERENCES contract(id)
	);`
	create_shift := `
	CREATE TABLE IF NOT EXISTS shift (
		id SERIAL PRIMARY KEY,
		employment_id INT,
		start_date TIMESTAMPTZ,
		end_date TIMESTAMPTZ,
		FOREIGN KEY (employment_id) REFERENCES employment(id)
	);`
	create_refresh_token := `
	CREATE TABLE IF NOT EXISTS refresh_token (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		device_id TEXT NOT NULL,
		token_hash TEXT NOT NULL,
		expires_at TIMESTAMP NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		UNIQUE(user_id, device_id)
	);`

	create(db, create_user)
	create(db, create_company)
	create(db, create_user_pin_auth)
	create(db, create_user_password_auth)
	create(db, create_contract)
	create(db, create_employment)
	create(db, create_shift)
	create(db, create_refresh_token)
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
