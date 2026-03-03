package auth

import "context"

type contextKey string

const ClaimsKey contextKey = "claims"
const DeviceIdKey contextKey = "deviceID"

func ClaimsFromContext(ctx context.Context) (*Claims, bool) {
	claims, ok := ctx.Value(ClaimsKey).(*Claims)
	return claims, ok
}
