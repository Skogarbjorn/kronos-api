package main

import (
	"database/sql"
	"log"
	"test/internal/router"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgresql://hbv2db_user:dawuJpPDAWjvUoYWSUNmuMLQZY2NHX98@dpg-d5npj9er433s739tg3eg-a.frankfurt-postgres.render.com/hbv2db"
	
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

	//api.DropTables(db)
	//api.CreateTables(db)

	r := router.CreateRouter(db)

	log.Fatal(router.RunServer(":8080", r))
}
