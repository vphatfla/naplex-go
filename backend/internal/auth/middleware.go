package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/vphatfla/naplex-go/backend/internal/config"
	"github.com/vphatfla/naplex-go/backend/internal/utils"
)

const (
	SESSION_CONTEXT_KEY = "contextKey"
)

type Middleware struct {
	cookieManager *CookieManager
}

func NewMiddleware(config *config.Config) *Middleware{
	return &Middleware{
		cookieManager: NewCookieManager(config.CookieSecret),
	}
}

func (m *Middleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie("session")
		if err != nil {
			utils.HTTPJsonError(w, fmt.Errorf("Error getting cookie -> %v", err).Error(), http.StatusBadRequest)
			return
		}

		var sData Session
		if err := m.cookieManager.ValidateCookie(sessionCookie, &sData); err != nil {
			utils.HTTPJsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		if time.Now().After(sData.ExpiresAt) {
			utils.HTTPJsonError(w, fmt.Errorf("Expired session").Error(), http.StatusUnauthorized)
			return
		}

		// ctx := context.WithValue(r.Context(), SESSION_CONTEXT_KEY, sData)
		ctx := context.WithValue(r.Context(), "user_id", sData.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getSessionDataFromContext(ctx context.Context) *Session {
	session, ok := ctx.Value(SESSION_CONTEXT_KEY).(*Session)
	if !ok {
		return nil
	}
	return session
}
