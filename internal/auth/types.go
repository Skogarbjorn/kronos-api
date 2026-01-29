package auth

type UserCreate struct {
	KT        string `json:"kt"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	Pin      *string `json:"pin,omitempty"`
}

type UserPinAuth struct {
	KT       string `json:"kt"`
	Pin      string `json:"pin"`
	DeviceID string `json:"device_id"`
}

type UserSilentRefresh struct {
	DeviceID string `json:"device_id"`
	RefreshToken string `json:"refresh_token"`
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

type UserReAuth struct {
	Pin          string `json:"pin"`
	RefreshToken string `json:"refresh_token"`
	DeviceID     string `json:"device_id"`
}

type UserPasswordAuth struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type User struct {
	ID        int    `json:"id"`
	KT        string `json:"kt"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AuthResponse struct {
	Message string `json:"message"`
	Tokens  Tokens `json:"tokens"`
	User    User   `json:"user"`
}
