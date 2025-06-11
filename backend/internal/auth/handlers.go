package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/vphatfla/naplex-go/backend/internal/config"
	"github.com/vphatfla/naplex-go/backend/internal/shared/database"
	"github.com/vphatfla/naplex-go/backend/internal/utils"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	config *config.Config
	cookieManager *CookieManager
	db *database.Queries
}

func NewAuthHandler(config *config.Config, db *database.Queries) *AuthHandler {
	return &AuthHandler{
		config: config,
		cookieManager: NewCookieManager(config.CookieSecret),
		db: db,
	}
}

func (h *AuthHandler) RegisterRouter() *http.ServeMux {
	m := http.NewServeMux()

	m.Handle("GET /google/login", http.HandlerFunc(h.HandleGoogleLogin))
	m.Handle("GET /google/callback", http.HandlerFunc(h.HandleGoogleCallback))

	return m
}
func (h *AuthHandler) HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	state, err := GenerateStateToken()
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

	/*utils.HTTPJsonResponse(w, map[string]string{
		"url": url,
	}) */

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *AuthHandler) HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("oauth_state")
	if err != nil {
		utils.HTTPJsonError(w, "missing cookie", http.StatusBadRequest)
		return
	}

	var stateData map[string]string
	if err := h.cookieManager.ValidateCookie(cookie, stateData); err != nil {
		utils.HTTPJsonError(w, "invalid state cookie", http.StatusBadRequest)
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

	googleUser, err := GetGoogleUserInfo(ctx, token)
	if err != nil {
		utils.HTTPJsonError(w, fmt.Errorf("Failed to retrieve Google User Info -> %v", err).Error(), http.StatusInternalServerError)
		return
	}

	params := &database.CreateOrUpsertUserParams{
		GoogleID: googleUser.ID,
		Email:  googleUser.Email,
		Name: googleUser.Name,
		FirstName: pgtype.Text{
			String: googleUser.FirstName,
			Valid: true,
		},
		LastName:pgtype.Text{
			String: googleUser.LastName,
			Valid: true,
		},
		Picture:pgtype.Text{
			String: googleUser.Picture,
			Valid: true,
		},
	}

	u, err := h.db.CreateOrUpsertUser(ctx, *params)
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

	utils.HTTPJsonResponse(w, map[string]interface{} {
		"user": map[string]string{
			"user_id": string(u.ID),
			"email": u.Email,
			"name": u.Name,
			"picture": u.Picture.String,
		},
	})
}

func (h *AuthHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
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
