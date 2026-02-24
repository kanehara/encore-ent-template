package api

import (
	"context"

	"encore.dev/middleware"
)

// contextKey is an unexported type for context keys in this package.
// Using a named type avoids collisions with keys from other packages.
type contextKey string

const sessionIDKey contextKey = "session_id"

//encore:middleware target=all
func (s *Service) SessionMiddleware(req middleware.Request, next middleware.Next) middleware.Response {
	sessionID := req.Data().Headers.Get("x-session-id")
	ctx := context.WithValue(req.Context(), sessionIDKey, sessionID)
	return next(req.WithContext(ctx))
}

func getSessionID(ctx context.Context) string {
	if sessionID, ok := ctx.Value(sessionIDKey).(string); ok {
		return sessionID
	}
	return ""
}
