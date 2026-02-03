package manage

import (
	"database/sql"
	"net/http"
	"test/internal/abstractions"
)

func CreateWorkspaceHandler(db *sql.DB) http.HandlerFunc {
	return abstractions.CreateJSONHandler(db, CreateWorkspace)
}

func CreateCompanyHandler(db *sql.DB) http.HandlerFunc {
	return abstractions.CreateJSONHandler(db, CreateCompany)
}

func CreateLocationHandler(db *sql.DB) http.HandlerFunc {
	return abstractions.CreateJSONHandler(db, CreateLocation)
}
