package manage

import (
	"database/sql"
	"encoding/json"
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

func GetWorkspacesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := GetWorkspaces(r.Context(), db)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func GetCompaniesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := GetCompanies(r.Context(), db)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func GetLocationsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := GetLocations(r.Context(), db)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func GetTasksHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := GetTasks(r.Context(), db)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func GetProfilesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := GetProfiles(r.Context(), db)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func GetEmploymentsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := GetEmployments(r.Context(), db)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func GetContractsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := GetContracts(r.Context(), db)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func GetShiftsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := GetShifts(r.Context(), db)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
