package handlers


import (
	httper "medods/http"
	"medods/services"
	"medods/utils"
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

	return map[string]string{"access_token": access, "refresh_token": refresh}, nil
}

func refresh(h httper.Helper) (any, error) {
	token, err := h.GetQueryParamAsString("token")
	if err != nil {
		return nil, err
	}

	claims, err := utils.VerificationToken(token)
	if err != nil {
		return nil, err
	}

	if float64(time.Now().Unix()) > claims["expire_at"].(float64) {
		return nil, httper.NewHttpErrorBadRequest("refresh token expired")
	}

	access, refresh, err := h.Service.(services.AuthService).TryRefreshToken(token, claims["uuid"].(string))
	if err != nil {
		return nil, err
	}

	return map[string]string{"access_token": access, "refresh_token": refresh}, nil
}
