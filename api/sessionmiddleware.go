package api

import (
	"context"

	"encore.dev/middleware"
)

//encore:middleware target=all
func (s *Service) SessionMiddleware(req middleware.Request, next middleware.Next) middleware.Response {
	sessionID := req.Data().Headers.Get("x-session-id")
	ctx := context.WithValue(req.Context(), "session_id", sessionID)
	return next(req.WithContext(ctx))
}

func getSessionID(ctx context.Context) string {
	if sessionID, ok := ctx.Value("session_id").(string); ok {
		return sessionID
	}
	return ""
}
