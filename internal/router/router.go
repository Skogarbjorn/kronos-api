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
	"github.com/go-chi/cors"
)

func CreateRouter(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://kronos-website-production-0d7c.up.railway.app", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Device-ID"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/checkhealth", checkhealthHandler(db))
		r.Route("/auth", func(r chi.Router) {
			r.Use(auth.DeviceIdMiddleware())

			r.Post("/register", auth.RegisterHandler(db))
			r.Post("/login", auth.LoginHandler(db))
			r.Post("/refresh", auth.SilentRefreshHandler(db))
			r.Post("/reauth", auth.ReAuthHandler(db))
		})

		r.Route("/manage", func(r chi.Router) {
			r.Post("/workspace",  manage.CreateWorkspaceHandler(db))
			r.Post("/company",    manage.CreateCompanyHandler(db))
			r.Post("/location",   manage.CreateLocationHandler(db))
			r.Post("/task",       manage.CreateTaskHandler(db))
			r.Post("/contract",   manage.CreateContractHandler(db))
			r.Post("/employment", manage.CreateEmploymentHandler(db))
			r.Post("/profile",    auth.RegisterHandler(db))

			r.Get("/workspaces",  manage.GetWorkspacesHandler(db))
			r.Get("/companies",   manage.GetCompaniesHandler(db))
			r.Get("/locations",   manage.GetLocationsHandler(db))
			r.Get("/tasks",       manage.GetTasksHandler(db))
			r.Get("/profiles",    manage.GetProfilesHandler(db))
			r.Get("/employments",    manage.GetEmploymentsHandler(db))
			r.Get("/contracts",    manage.GetContractsHandler(db))
			r.Get("/shifts",      manage.GetShiftsHandler(db))

			r.Delete("/workspaces/{id}",  manage.DeleteWorkspaceHandler(db))
			r.Delete("/companies/{id}",   manage.DeleteCompanyHandler(db))
			r.Delete("/locations/{id}",   manage.DeleteLocationHandler(db))
			r.Delete("/tasks/{id}",       manage.DeleteTaskHandler(db))
			r.Delete("/profiles/{id}",    manage.DeleteProfileHandler(db))
			r.Delete("/contracts/{id}",   manage.DeleteContractHandler(db))
			r.Delete("/shifts/{id}",      manage.DeleteShiftHandler(db))

			r.Patch("/workspaces/{id}",  manage.PatchWorkspaceHandler(db))
			r.Patch("/companies/{id}",   manage.PatchCompanyHandler(db))
			r.Patch("/locations/{id}",   manage.PatchLocationHandler(db))
			r.Patch("/tasks/{id}",       manage.PatchTaskHandler(db))
			r.Patch("/profiles/{id}",    manage.PatchProfileHandler(db))
			r.Patch("/contracts/{id}",   manage.PatchContractHandler(db))
			r.Patch("/shifts/{id}",      manage.PatchShiftHandler(db))
		})

		r.Route("/pin", func(r chi.Router) {
			r.Use(auth.PinAuthMiddleware([]byte(os.Getenv("JWT_SECRET"))))
			r.Use(auth.DeviceIdMiddleware())

			r.Post("/clock-in", pin.ClockInHandler(db))
			r.Post("/clock-out", pin.ClockOutHandler(db))
			r.Post("/sync-shift", pin.SyncShiftHandler(db))
			r.Get("/shift-overview", pin.ShiftOverviewHandler(db))
			r.Get("/shift-history", pin.ShiftHistoryHandler(db))
			r.Get("/locations", pin.GetLocationsHandler(db))
			r.Get("/tasks", pin.GetTasksHandler(db))
			r.Get("/employments-detailed", pin.GetEmploymentsDetailedHandler(db))
			r.Get("/my-pin", pin.GetPinHandler(db))
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
