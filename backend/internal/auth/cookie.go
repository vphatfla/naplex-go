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
	"strings"
	"time"
)

// Data encrypted and stored in the cookie
type Session struct {
	// UserID field id in the user table is SERIAL, which is 4 bytes -> int32 in golang
	UserID int32`json:"user_id"`
	Email string `json:"email"`
	Name string `json:"name"`
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

// Validate cookie
// dest is a flexible data type that can hold pointer as value
func (cm *CookieManager) ValidateCookie(cookie *http.Cookie, dest interface{}) error {
	parts := strings.Split(cookie.Value, ".")
	if len(parts) != 2 {
		return fmt.Errorf("Invalid cookie length")
	}

	// decode
	eData, err := base64.URLEncoding.DecodeString(parts[0])
	if err != nil {
		return fmt.Errorf("ValidateCookie: error decoding session data -> %v", err)
	}

	signature, err := base64.URLEncoding.DecodeString(parts[1])
	if err != nil {
		return fmt.Errorf("ValidateCookie: error decoding signature -> %v", err)
	}
	expectedSignature := cm.sign(eData)
	if !hmac.Equal(signature, expectedSignature) {
		return fmt.Errorf("ValidateCookie: Invalid cookie, not equal")
	}

	decryptData, err := cm.decrypt(eData)
	if err != nil {
		return fmt.Errorf("ValidateCookie: error decrypting session -> %v", err)
	}

	if err := json.Unmarshal(decryptData, &dest); err != nil {
		return fmt.Errorf("ValidateCookie: error unmarshaling data -> %v", err)
	}

	return nil
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
func (cm *CookieManager) decrypt(data []byte) ([]byte, error) {
    block, err := aes.NewCipher(cm.cookieSecret)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    if len(data) < gcm.NonceSize() {
        return nil, fmt.Errorf("ciphertext too short")
    }

    nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
    return gcm.Open(nil, nonce, ciphertext, nil)
}
