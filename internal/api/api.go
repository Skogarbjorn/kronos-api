package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"test/internal/lib"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateRouter(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/checkhealth", checkhealthHandler(db))
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", registerHandler(db))
			r.Post("/login", loginHandler(db))
			r.Post("/refresh", silentRefreshHandler(db))
			r.Post("/reauth", reAuthHandler(db))
		})
	})

	return r
}

func RunServer(addr string, handler http.Handler) error {
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started on %s", addr)

	return server.ListenAndServe()
}

func registerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input lib.UserCreate
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}

		user, err := lib.CreateUser(r.Context(), db, input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func loginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input lib.UserPinAuth
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}

		res, err := lib.ColdStartPin(r.Context(), db, input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(res)
	}
}

func silentRefreshHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input lib.UserSilentRefresh
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}

		tokens, err := lib.RefreshTokens(r.Context(), db, input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(tokens)
	}
}

func reAuthHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input lib.UserReAuth
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}

		tokens, err := lib.WarmStartPin(r.Context(), db, input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(tokens)
	}
}

func checkhealthHandler(_ *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}
}
