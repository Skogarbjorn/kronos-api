package main

import (
	"database/sql"
	"log"
	"os"
	"test/internal/router"

	//dbrepo "test/internal/db"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file %w", err)
	}

	connStr := os.Getenv("DATABASE_CONNECTION_STRING")
	
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Connected to PostgreSQL successfully!")

	//dbrepo.DropTables(db)
	//dbrepo.CreateTables(db)
	//dbrepo.InsertDummy(db)

	r := router.CreateRouter(db)

	port := ":" + os.Getenv("PORT")
	log.Fatal(router.RunServer(port, r))
}
