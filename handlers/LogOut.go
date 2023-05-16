package handlers

import (
	"net/http"
)

func LogOut(w http.ResponseWriter, r *http.Request) {
	tokenCookie := http.Cookie{
		Name:     "Token",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}
	http.SetCookie(w, &tokenCookie)
	refreshCookie := http.Cookie{
		Name:     "Refresh",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}
	http.SetCookie(w, &refreshCookie)
	w.WriteHeader(http.StatusOK)
}
