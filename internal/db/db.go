package db

import (
	"database/sql"
	"log"
	"os"
)

func insert(db *sql.DB, row string) {
	_, err := db.Exec(row)
	if err != nil {
		log.Fatalf("Failed to insert row: %v", err)
	}
	log.Println("Row inserted successfully.")
}

func CreateTables(db *sql.DB) {
	b, err := os.ReadFile("internal/db/create.sql")
	_, err = db.Exec(string(b))
	if err != nil {
		log.Fatalf("Failed create db actions: %v", err)
	}
	log.Println("Create db schema successful")
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
