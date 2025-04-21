package handlers

import (
	httper "medods/http"
	"medods/services"
	"medods/utils"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func BindAuth(router *chi.Mux) {
	router.Get("/login", httper.CreateHandler(login, services.Auth()))
	router.Get("/refresh", httper.CreateHandler(refresh, services.Auth()))
}

func login(h httper.Helper) (any, error) {
	uuid, err := h.GetQueryParamAsString("uuid")
	if err != nil {
		return nil, err
	}

	access, refresh, err := h.Service.(services.AuthService).CreateTokens(uuid)
	if err != nil {
		return nil, err
	}

	http.SetCookie(h.W, &http.Cookie{
		Name:     "refresh_token",
		Value:    refresh,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	return map[string]string{"access_token": access, "refresh_token": refresh}, nil
}

func refresh(h httper.Helper) (any, error) {
	token, err := h.R.Cookie("refresh_token")
	if err != nil {
		return nil, err
	}

	claims, err := utils.VerificationToken(token.Value)
	if err != nil {
		return nil, err
	}

	if float64(time.Now().Unix()) > claims["expire_at"].(float64) {
		return nil, httper.NewHttpErrorBadRequest("refresh token expired")
	}

	access, refresh, err := h.Service.(services.AuthService).TryRefreshToken(token.Value, claims["uuid"].(string))
	if err != nil {
		return nil, err
	}

	http.SetCookie(h.W, &http.Cookie{
		Name:     "refresh_token",
		Value:    refresh,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	return map[string]string{"access_token": access, "refresh_token": refresh}, nil
}
