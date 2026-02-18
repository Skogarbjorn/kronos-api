package auth

import "github.com/golang-jwt/jwt/v5"

type ProfileCreate struct {
	KT        string `json:"kt"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	Pin      *string `json:"pin,omitempty"`
}

type ProfileSilentRefresh struct {
	DeviceID string `json:"device_id"`
	RefreshToken string `json:"refresh_token"`
}

type ProfilePinAuth struct {
	KT       string `json:"kt"`
	Pin      string `json:"pin"`
	DeviceID string `json:"device_id"`
}

type ProfileReAuth struct {
	Pin          string `json:"pin"`
	RefreshToken string `json:"refresh_token"`
	DeviceID     string `json:"device_id"`
}

type ProfilePasswordAuth struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Tokens struct {
	AccessToken AccessToken `json:"access_token"`
	RefreshToken RefreshToken `json:"refresh_token"`
}

type RefreshToken struct {
	Token string `json:"token"`
	ExpiresAt int64 `json:"expires_at"`
}

type AccessToken struct {
	Token string `json:"token"`
	ExpiresAt int64 `json:"expires_at"`
}

type Profile struct {
	ID        int    `json:"id"`
	KT        string `json:"kt"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AuthResponse struct {
	Message string `json:"message"`
	Tokens  Tokens `json:"tokens"`
	Profile    Profile   `json:"profile"`
}

type Claims struct {
	ProfileID int    `json:"sub"`
	Auth   string `json:"auth"`
	jwt.RegisteredClaims
}
