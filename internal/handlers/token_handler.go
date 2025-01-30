package handlers

import (
	"encoding/json"
	"github.com/eenees/twitch-genie-server/internal/services"
	"io"
	"net/http"
)

type TokenHandler struct {
	service *services.TokenService
}

func NewTokenHandler(service *services.TokenService) *TokenHandler {
	return &TokenHandler{service: service}
}

type exchangeTokenBody struct {
	Code string `json:"code"`
}

// ExchangeToken godoc
//
// @Summary Exchange token
// @Description Exchange the auth token and retrieve user data
// @Accepts json
// @Produce json
// @Param exchangeTokenBody body exchangeTokenBody true "Exchange token body"
// @router /exchange-token [post]
// @Security ApiKeyAuth
// @Tags Authentication
func (handler *TokenHandler) ExchangeToken(w http.ResponseWriter, r *http.Request) {

	var body struct {
		Code string `json:"code" binding:"required"`
	}

	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(rawBody, &body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if body.Code == "" {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}

	tokenData, err := handler.service.ExchangeToken(body.Code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userData, err := handler.service.ValidateToken(tokenData.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.service.SaveUser(userData.UserId, userData.Login, tokenData.AccessToken, tokenData.RefreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "jwt value",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode, // TODO: change this
		Path:     "/",
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userData)
}
