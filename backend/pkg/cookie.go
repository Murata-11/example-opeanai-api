package pkg

import (
	"net/http"
	"time"
)

func SetAccessToken(token string) *http.Cookie {
	return &http.Cookie{
		Name:     "AccessToken",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(1 * time.Minute),
		Path:     "/",
	}
}

func SetRefreshToken(token string) *http.Cookie {
	return &http.Cookie{
		Name:     "RefreshToken",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
	}
}
