package router

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"test/internal/auth"
	"test/internal/manage"
	"test/internal/pin"
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
			r.Post("/register", auth.RegisterHandler(db))
			r.Post("/login", auth.LoginHandler(db))
			r.Post("/refresh", auth.SilentRefreshHandler(db))
			r.Post("/reauth", auth.ReAuthHandler(db))
		})

		r.Route("/manage", func(r chi.Router) {
			r.Post("/workspace", manage.CreateWorkspaceHandler(db))
			r.Post("/company", manage.CreateCompanyHandler(db))
			r.Post("/location", manage.CreateLocationHandler(db))
			r.Post("/task", manage.CreateTaskHandler(db))
			r.Post("/contract", manage.CreateContractHandler(db))
			r.Post("/employment", manage.CreateEmploymentHandler(db))
		})

		r.Route("/pin", func(r chi.Router) {
			r.Use(auth.PinAuthMiddleware([]byte(os.Getenv("JWT_SECRET"))))

			r.Post("/clock-in", pin.ClockInHandler(db))
			r.Post("/clock-out", pin.ClockOutHandler(db))
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

func checkhealthHandler(_ *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}
}
