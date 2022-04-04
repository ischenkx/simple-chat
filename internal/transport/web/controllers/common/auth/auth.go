package auth

import (
	"net/http"
)

func LoadVerificationToken(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func DeleteVerificationToken(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Secure:   true,
		HttpOnly: true,
	})
}

func StoreVerificationToken(w http.ResponseWriter, r *http.Request, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	})
}
