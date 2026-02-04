package manage

import (
	"database/sql"
	"net/http"
	"test/internal/abstractions"
)

func CreateWorkspaceHandler(db *sql.DB) http.HandlerFunc {
	return abstractions.CreateJSONHandler(db, CreateWorkspace, WriteDomainError)
}

func CreateCompanyHandler(db *sql.DB) http.HandlerFunc {
	return abstractions.CreateJSONHandler(db, CreateCompany, WriteDomainError)
}

func CreateLocationHandler(db *sql.DB) http.HandlerFunc {
	return abstractions.CreateJSONHandler(db, CreateLocation, WriteDomainError)
}

func CreateTaskHandler(db *sql.DB) http.HandlerFunc {
	return abstractions.CreateJSONHandler(db, CreateTask, WriteDomainError)
}

func CreateEmploymentHandler(db *sql.DB) http.HandlerFunc {
	return abstractions.CreateJSONHandler(db, CreateEmployment, WriteDomainError)
}

func CreateContractHandler(db *sql.DB) http.HandlerFunc {
	return abstractions.CreateJSONHandler(db, CreateContract, WriteDomainError)
}
