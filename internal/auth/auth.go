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
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(
	ctx context.Context,
	db *sql.DB,
	input UserCreate,
) (*User, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("CreateUser: begin tx: %w", err)
	}
	defer tx.Rollback()

	var user User
	err = tx.QueryRowContext(
		ctx,
		`
		INSERT INTO users (kt, first_name, last_name)
		VALUES ($1, $2, $3)
		RETURNING id, kt, first_name, last_name
		`,
		input.KT,
		input.FirstName,
		input.LastName,
	).Scan(
		&user.ID,
		&user.KT,
		&user.FirstName,
		&user.LastName,
	)
	if err != nil {
		return nil, fmt.Errorf("CreateUser: db insert: %w", err)
	}

	if input.Pin != nil {
		err = addPinAuth(ctx, tx, user.ID, *input.Pin)
		if err != nil {
			return nil, fmt.Errorf("CreateUser: %w", err)
		}
	}
	if input.Password != nil {
		err = addPasswordAuth(ctx, tx, user.ID, *input.Password, *input.Email)
		if err != nil {
			return nil, fmt.Errorf("CreateUser: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("CreateUser: db commit: %w", err)
	}

	return &user, nil
}

func addPinAuth(
	ctx context.Context,
	tx *sql.Tx,
	user_id int,
	pin string,
) error {
	pinHash := hashPin(pin, "todo!")
	_, err := tx.ExecContext(
		ctx,
		`
		INSERT INTO user_pin_auth (user_id, pin)
		VALUES ($1, $2)
		`,
		user_id,
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
	user_id int,
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
		INSERT INTO user_password_auth (user_id, email, password)
		VALUES ($1, $2, $3)
		`,
		user_id,
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
	input UserPinAuth,
) (*AuthResponse, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("AuthenticateUser: begin tx: %w", err)
	}
	defer tx.Rollback()

	var (
		user    User
		pinHash string
	)

	err = tx.QueryRowContext(
		ctx,
		`
		SELECT
			u.id,
			u.kt,
			u.first_name,
			u.last_name,
			p.pin
		FROM users u
		JOIN user_pin_auth p ON p.user_id = u.id
		WHERE u.kt = $1
		`,
		input.KT,
	).Scan(
		&user.ID,
		&user.KT,
		&user.FirstName,
		&user.LastName,
		&pinHash,
	)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}; if err != nil {
		return nil, fmt.Errorf("AuthenticateUser: query user: %w", err)
	}

	inputPinHash := hashPin(input.Pin, "todo!")
	if subtle.ConstantTimeCompare(
		[]byte(pinHash),
		[]byte(inputPinHash),
	) != 1 {
		return nil, ErrInvalidCredentials
	}

	accessToken, err := createAccessToken(user.ID)
	if err != nil {
		return nil, err
	}
	refreshToken, err := createRefreshToken(ctx, tx, user.ID, input.DeviceID)
	if err != nil {
		return nil, err
	}

	response := AuthResponse{
		Message: "Login successful",
		User: user,
		Tokens: Tokens{
			AccessToken: *accessToken,
			RefreshToken: *refreshToken,
		},
	}

	tx.Commit()

	return &response, nil
}

func RefreshTokens(
	ctx context.Context,
	db *sql.DB,
	input UserSilentRefresh,
) (*Tokens, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("RefreshTokens: begin tx: %w", err)
	}
	defer tx.Rollback()

	hash := hashToken(input.RefreshToken)

	var user_id int
	var token_id int
	err = tx.QueryRowContext(
		ctx,
		`
		SELECT id, user_id FROM refresh_token WHERE 
		token_hash = $1 AND device_id = $2 AND expires_at > now()
		`,
		hash,
		input.DeviceID,
	).Scan(
		&token_id,
		&user_id,
	)
	if err != nil {
		return nil, fmt.Errorf("RefreshTokens: db select: %w", err)
	}

	access, refresh, err := rotateTokens(ctx, tx, user_id, input.DeviceID, token_id)
	if err != nil {
		return nil, fmt.Errorf("RefreshTokens: %w", err)
	}

	tx.Commit()

	tokens := Tokens{
		AccessToken: *access,
		RefreshToken: *refresh,
	}

	return &tokens, nil
}

func WarmStartPin(
	ctx context.Context,
	db *sql.DB,
	input UserReAuth,
) (*Tokens, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("WarmStartPin: begin tx: %w", err)
	}
	defer tx.Rollback()

	hashedToken := hashToken(input.RefreshToken)

	var (
		user_id int
		pinHash string
		token_id int
	)

	err = tx.QueryRowContext(
		ctx,
		`
		SELECT r.id, u.id, p.pin 
        FROM users u
        JOIN refresh_token r ON r.user_id = u.id
        JOIN user_pin_auth p ON p.user_id = u.id
        WHERE r.token_hash = $1 AND r.device_id = $2
		`,
		hashedToken,
		input.DeviceID,
	).Scan(
		&token_id,
		&user_id,
		&pinHash,
	)
	if err != nil {
		return nil, fmt.Errorf("WarmStartPin: db select: %w", err)
	}

	inputPinHash := hashPin(input.Pin, "todo!")
	if subtle.ConstantTimeCompare(
		[]byte(pinHash),
		[]byte(inputPinHash),
	) != 1 {
		return nil, ErrInvalidCredentials
	}

	accessToken, refreshToken, err := rotateTokens(ctx, tx, user_id, input.DeviceID, token_id)
	if err != nil {
		return nil, fmt.Errorf("WarmStartPin: %w", err)
	}

	tx.Commit()

	tokens := Tokens{
		AccessToken: *accessToken,
		RefreshToken: *refreshToken,
	}

	return &tokens, nil
}

func rotateTokens(
	ctx context.Context,
	tx *sql.Tx,
	user_id int,
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

	access, _ := createAccessToken(user_id)
	refresh, _ := createRefreshToken(ctx, tx, user_id, device_id)
	return access, refresh, nil
}

func createAccessToken(
	user_id int,
) (*AccessToken, error) {
	expiresAt := time.Now().Add(time.Hour).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": user_id,
			"exp": 	expiresAt,
		})

	accessTokenString, err := accessToken.SignedString([]byte("todo! create env and set secret"))
	if err != nil {
		return nil, fmt.Errorf("createAccessToken: sign jwt: %w", err)
	}

	return &AccessToken{ Token: accessTokenString, ExpiresAt: expiresAt }, nil
}

func createRefreshToken(
	ctx context.Context,
	tx *sql.Tx,
	user_id int,
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
		INSERT INTO refresh_token (user_id, device_id, token_hash, expires_at)
		VALUES ($1, $2, $3, now() + interval '12 hours')
		ON CONFLICT (user_id, device_id)
		DO UPDATE SET
			token_hash = EXCLUDED.token_hash,
			expires_at = EXCLUDED.expires_at,
			created_at = now()
		`,
		user_id,
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
