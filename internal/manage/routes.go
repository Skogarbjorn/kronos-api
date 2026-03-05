package manage

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"test/internal/abstractions"

	"github.com/go-chi/chi/v5"
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

func DeleteWorkspaceHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		result, err := DeleteWorkspace(r.Context(), db, id)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func DeleteCompanyHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		result, err := DeleteCompany(r.Context(), db, id)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func DeleteLocationHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		result, err := DeleteLocation(r.Context(), db, id)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func DeleteTaskHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		result, err := DeleteTask(r.Context(), db, id)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func DeleteProfileHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		result, err := DeleteProfile(r.Context(), db, id)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func DeleteContractHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		result, err := DeleteContract(r.Context(), db, id)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func DeleteShiftHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		result, err := DeleteShift(r.Context(), db, id)
		if err != nil {
			WriteDomainError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
