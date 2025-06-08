package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Data encrypted and stored in the cookie
type Session struct {
	UserID string `json:"user_id"`
	Email string `json:"user_id"`
	Name string `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}

// Cookie manager handle all the operations of cookie
type CookieManager struct {
	cookieSecret []byte
}

func NewCookieManager(secret []byte) *CookieManager {
	return &CookieManager{
		cookieSecret: secret,
	}
}

// Encrypt and sign the cookie
func (cm *CookieManager) CreateCookie(name string, data interface{}, maxAge int) (*http.Cookie, error) {
	bData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("CreateCookie: Error marshalling data -> %v", err)
	}

	eData, err := cm.encrypt(bData)
	if err != nil {
		return nil, fmt.Errorf("CreatCookie: error encrypt data -> %v", err)
	}

	signature := cm.sign(eData)

	value := base64.URLEncoding.EncodeToString(eData) + "." + base64.URLEncoding.EncodeToString(signature)

	return &http.Cookie{
		Name: name,
		Value: value,
		Path:     "/",
        MaxAge:   maxAge,
        HttpOnly: true,
        Secure:   false, // !cm.isDevelopment, // Use Secure in production
        SameSite: http.SameSiteLaxMode,
    }, nil
}

func (cm *CookieManager) encrypt(data []byte) ([]byte, error) {
    block, err := aes.NewCipher(cm.cookieSecret)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }

    return gcm.Seal(nonce, nonce, data, nil), nil
}

func (cm *CookieManager) sign(data []byte) []byte {
	h := hmac.New(sha256.New, cm.cookieSecret)
	h.Write(data)
	return h.Sum(nil)
}
