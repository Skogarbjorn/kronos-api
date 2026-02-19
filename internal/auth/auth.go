package auth

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateProfile(
	ctx context.Context,
	db *sql.DB,
	input ProfileCreate,
) (*Profile, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("CreateProfile: begin tx: %w", err)
	}
	defer tx.Rollback()

	var profile Profile
	err = tx.QueryRowContext(
		ctx,
		`
		INSERT INTO profile (kt, first_name, last_name)
		VALUES ($1, $2, $3)
		RETURNING id, kt, first_name, last_name
		`,
		input.KT,
		input.FirstName,
		input.LastName,
	).Scan(
		&profile.ID,
		&profile.KT,
		&profile.FirstName,
		&profile.LastName,
	)
	if err != nil {
		return nil, fmt.Errorf("CreateProfile: db insert: %w", err)
	}

	if input.Pin != nil {
		err = addPinAuth(ctx, tx, profile.ID, *input.Pin)
		if err != nil {
			return nil, fmt.Errorf("CreateProfile: %w", err)
		}
	}
	if input.Password != nil {
		err = addPasswordAuth(ctx, tx, profile.ID, *input.Password, *input.Email)
		if err != nil {
			return nil, fmt.Errorf("CreateProfile: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("CreateProfile: db commit: %w", err)
	}

	return &profile, nil
}

func addPinAuth(
	ctx context.Context,
	tx *sql.Tx,
	profile_id int,
	pin string,
) error {
	pinHash := hashPin(pin, "todo!")
	_, err := tx.ExecContext(
		ctx,
		`
		INSERT INTO profile_pin_auth (profile_id, pin)
		VALUES ($1, $2)
		`,
		profile_id,
		pinHash,
	)
	if err != nil {
		return fmt.Errorf("addPinAuth: %w", err)
	}
	return nil
}

func hashPin(pin string, _ string) string {
	h := hmac.New(sha256.New, []byte("todo!"))
	h.Write([]byte(pin))
	return hex.EncodeToString(h.Sum(nil))
}

func addPasswordAuth(
	ctx context.Context,
	tx *sql.Tx,
	profile_id int,
	password string,
	email string,
) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("addPasswordAuth: hash: %w", err)
	}

	_, err = tx.ExecContext(
		ctx,
		`
		INSERT INTO profile_password_auth (profile_id, email, password)
		VALUES ($1, $2, $3)
		`,
		profile_id,
		email,
		passwordHash,
	)

	if err != nil {
		return fmt.Errorf("addPasswordAuth: db insert: %w", err)
	}

	return nil
}

func ColdStartPin(
	ctx context.Context,
	db *sql.DB,
	input ProfilePinAuth,
) (*AuthResponse, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("AuthenticateProfile: begin tx: %w", err)
	}
	defer tx.Rollback()
	println("after making tx")

	var (
		profile    Profile
		pinHash string
	)

	println("after init profile + pinHash")

	println("before making query")
	err = tx.QueryRowContext(
		ctx,
		`
		SELECT
			u.id,
			u.kt,
			u.first_name,
			u.last_name,
			p.pin
		FROM profile u
		JOIN profile_pin_auth p ON p.profile_id = u.id
		WHERE u.kt = $1
		`,
		input.KT,
	).Scan(
		&profile.ID,
		&profile.KT,
		&profile.FirstName,
		&profile.LastName,
		&pinHash,
	)
	println("before handling sql errors")
	if err == sql.ErrNoRows {
		return nil, ErrProfileNotFound
	}; if err != nil {
		return nil, fmt.Errorf("AuthenticateProfile: query profile: %w", err)
	}

	println("before hashing")

	inputPinHash := hashPin(input.Pin, "todo!")
	if subtle.ConstantTimeCompare(
		[]byte(pinHash),
		[]byte(inputPinHash),
	) != 1 {
		return nil, ErrInvalidCredentials
	}

	println("before creating tokens")

	accessToken, err := createAccessToken(profile.ID, "pin")
	if err != nil {
		return nil, err
	}
	refreshToken, err := createRefreshToken(ctx, tx, profile.ID, input.DeviceID)
	if err != nil {
		return nil, err
	}

	println("before creating response")

	response := AuthResponse{
		Message: "Login successful",
		Profile: profile,
		Tokens: Tokens{
			AccessToken: *accessToken,
			RefreshToken: *refreshToken,
		},
	}
	println(response.Message)

	tx.Commit()

	return &response, nil
}

func RefreshTokens(
	ctx context.Context,
	db *sql.DB,
	input ProfileSilentRefresh,
) (*AuthResponse, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("RefreshTokens: begin tx: %w", err)
	}
	defer tx.Rollback()

	hash := hashToken(input.RefreshToken)

	var (
	    profile_id int
	    token_id int
		profile Profile
	)
	err = tx.QueryRowContext(
		ctx,
		`
		SELECT r.id, r.profile_id, p.id, p.kt, p.first_name, p.last_name
		FROM refresh_token r
		JOIN profile p ON p.id = r.profile_id
		WHERE r.token_hash = $1 AND r.device_id = $2 AND r.expires_at > now()
		`,
		hash,
		input.DeviceID,
	).Scan(
		&token_id,
		&profile_id,
	)
	if err != nil {
		return nil, fmt.Errorf("RefreshTokens: db select: %w", err)
	}

	access, refresh, err := rotateTokens(ctx, tx, profile_id, input.DeviceID, token_id)
	if err != nil {
		return nil, fmt.Errorf("RefreshTokens: %w", err)
	}

	tx.Commit()

	tokens := Tokens{
		AccessToken: *access,
		RefreshToken: *refresh,
	}

	response := AuthResponse{
		Message: "Silent refresh successful",
		Tokens: tokens,
		Profile: profile,
	}

	return &response, nil
}

func WarmStartPin(
	ctx context.Context,
	db *sql.DB,
	input ProfileReAuth,
) (*AuthResponse, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("WarmStartPin: begin tx: %w", err)
	}
	defer tx.Rollback()

	hashedToken := hashToken(input.RefreshToken)

	var (
		profile_id int
		pinHash string
		token_id int
		profile Profile
	)

	err = tx.QueryRowContext(
		ctx,
		`
		SELECT r.id, u.id, p.pin, u.id, u.kt, u.first_name, u.last_name
        FROM profile u
        JOIN refresh_token r ON r.profile_id = u.id
        JOIN profile_pin_auth p ON p.profile_id = u.id
        WHERE r.token_hash = $1 AND r.device_id = $2
		`,
		hashedToken,
		input.DeviceID,
	).Scan(
		&token_id,
		&profile_id,
		&pinHash,
		&profile.ID,
		&profile.KT,
		&profile.FirstName,
		&profile.LastName,
	)
	if err != nil {
		return nil, fmt.Errorf("WarmStartPin: db select: %w", err)
	}

	inputPinHash := hashPin(input.Pin, os.Getenv("PIN_HASH_SECRET"))
	if subtle.ConstantTimeCompare(
		[]byte(pinHash),
		[]byte(inputPinHash),
	) != 1 {
		return nil, ErrInvalidCredentials
	}

	accessToken, refreshToken, err := rotateTokens(ctx, tx, profile_id, input.DeviceID, token_id)
	if err != nil {
		return nil, fmt.Errorf("WarmStartPin: %w", err)
	}

	tx.Commit()

	tokens := Tokens{
		AccessToken: *accessToken,
		RefreshToken: *refreshToken,
	}

	response := AuthResponse{
		Message: "Authentication successful",
		Tokens: tokens,
		Profile: profile,
	}

	return &response, nil
}

func rotateTokens(
	ctx context.Context,
	tx *sql.Tx,
	profile_id int,
	device_id string,
	old_token_id int,
) (*AccessToken, *RefreshToken, error) {
	_, err := tx.ExecContext(
		ctx,
		`
		DELETE FROM refresh_token WHERE id = $1
		`,
		old_token_id,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("rotateTokens: %w", err)
	}

	access, _ := createAccessToken(profile_id, "pin")
	refresh, _ := createRefreshToken(ctx, tx, profile_id, device_id)
	return access, refresh, nil
}

func createAccessToken(
	profile_id int,
	auth string,
) (*AccessToken, error) {
	expiresAt := time.Now().Add(time.Hour)
	claims := Claims{
		ProfileID: profile_id,
		Auth: auth,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, fmt.Errorf("createAccessToken: sign jwt: %w", err)
	}

	return &AccessToken{ Token: accessTokenString, ExpiresAt: expiresAt.Unix() }, nil
}

func createRefreshToken(
	ctx context.Context,
	tx *sql.Tx,
	profile_id int,
	device_id string,
) (*RefreshToken, error) {
	token, err := generateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("createRefreshToken: gen: %w", err)
	}

	tokenHash := hashToken(token)

	_, err = tx.ExecContext(
		ctx,
		`
		INSERT INTO refresh_token (profile_id, device_id, token_hash, expires_at)
		VALUES ($1, $2, $3, now() + interval '12 hours')
		ON CONFLICT (profile_id, device_id)
		DO UPDATE SET
			token_hash = EXCLUDED.token_hash,
			expires_at = EXCLUDED.expires_at,
			created_at = now()
		`,
		profile_id,
		device_id,
		tokenHash,
	)
	if err != nil {
		return nil, fmt.Errorf("createRefreshToken: db insert: %w", err)
	}

	expiresAt := time.Now().Add(time.Hour * 12).Unix()
	return &RefreshToken{ Token: token, ExpiresAt: expiresAt }, nil
}

func generateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("generateRefreshToken: rand: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func hashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

func PinAuthMiddleware(secret []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(header, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			tokenStr := parts[1]

			claims := &Claims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return secret, nil
			})

			if err != nil || !token.Valid {
				print(err)
				print(token.Valid)
				print(token)
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			if claims.Auth != "pin" {
				http.Error(w, "invalid auth stage", http.StatusForbidden)
				return
			}
			
			ctx := context.WithValue(r.Context(), ClaimsKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
