package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/vphatfla/naplex-go/backend/internal/config"
	"github.com/vphatfla/naplex-go/backend/internal/utils"
	"golang.org/x/oauth2"
)

type Handler struct {
	config *config.Config
	cookieManager *CookieManager
	s *Service
}

func NewHandler(config *config.Config, s *Service) *Handler {
	return &Handler{
		config: config,
		cookieManager: NewCookieManager(config.CookieSecret),
		s: s,
	}
}

/* func (h *Handler) RegisterRouter() *http.ServeMux {
	m := http.NewServeMux()

	m.Handle("GET /google/login", http.HandlerFunc(h.HandleGoogleLogin))
	m.Handle("GET /google/callback", http.HandlerFunc(h.HandleGoogleCallback))

	return m
}*/
func (h *Handler) HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	state, err := h.s.GenerateStateToken()
	if err != nil {
		utils.HTTPJsonError(w, "Failed to generate state for session", http.StatusInternalServerError)
		return
	}

	cookie, err := h.cookieManager.CreateCookie("oauth_state", map[string]string{
        "state": state,
    }, 600) // 10 minutes)

	if err != nil {
		utils.HTTPJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, cookie)

	url := h.config.OAuth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline)

	log.Printf("Redirect uri -> %s", h.config.OAuth2Config.RedirectURL)
	log.Printf("URL -> %s", url)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("oauth_state")
	if err != nil {
		utils.HTTPJsonError(w, "missing cookie", http.StatusBadRequest)
		return
	}

	var stateData map[string]interface{}
	if err := h.cookieManager.ValidateCookie(cookie, &stateData); err != nil {
		utils.HTTPJsonError(w, fmt.Errorf("invalid state cookie -> %v").Error(), http.StatusBadRequest)
		return
	}

	if r.FormValue("state") != stateData["state"] {
		utils.HTTPJsonError(w, "invalid state token receipt", http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
        Name:     "oauth_state",
        Value:    "",
        Path:     "/",
        MaxAge:   -1,
        HttpOnly: true,
    })

	code := r.FormValue("code")
	if code == "" {
		utils.HTTPJsonError(w, "Missing auth code", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	token, err := h.config.OAuth2Config.Exchange(ctx, code)
	if err != nil {
		utils.HTTPJsonError(w, "Failed to exchange code for token from auth provider", http.StatusInternalServerError)
		return
	}

	googleUser, err := h.s.GetGoogleUserInfo(ctx, token)
	if err != nil {
		utils.HTTPJsonError(w, fmt.Errorf("Failed to retrieve Google User Info -> %v", err).Error(), http.StatusInternalServerError)
		return
	}

	u, err := h.s.CreateOrUpdateUser(ctx, googleUser)
	if err != nil {
		utils.HTTPJsonError(w, fmt.Errorf("Failed to create or update user record -> %v", err).Error(), http.StatusInternalServerError)
		return
	}

	// create session data
	session := Session{
		UserID: int(u.ID),
		Email: u.Email,
		Name: u.Name,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour), // 7 days
	}

	cookie, err = h.cookieManager.CreateCookie("session", session, 86400*7) // 7 days
	if err != nil {
		utils.HTTPJsonError(w, fmt.Errorf("Failed to create cookie from session -> %v", err).Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/users/info", http.StatusPermanentRedirect)
}

func (h *Handler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
        Name:     "session",
        Value:    "",
        Path:     "/",
        MaxAge:   -1,
        HttpOnly: true,
    })
	utils.HTTPJsonResponse(w, map[string]bool{
		"success": true,
	})
}
