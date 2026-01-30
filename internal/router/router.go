package router

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"test/internal/auth"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-jwt/jwt/v5"
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
			r.Post("/authTest", pinAuthTestHandler(db))
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

func pinAuthTestHandler(_ *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "missing Authorization header", http.StatusBadRequest)
			return
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid Authorization header format", http.StatusBadRequest)
			return
		}

		token := parts[1]

		claims, err := parseToken(token, []byte("todo! create env and set secret"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(claims.Auth))
	}
}

func parseToken(tokenStr string, secret []byte) (*auth.Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&auth.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return secret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*auth.Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
