package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eenees/twitch-genie-server/internal/repositories"
	"github.com/eenees/twitch-genie-server/internal/utils/auth"
	"io"
	"net/http"
	"os"
)

type TokenService struct {
	repo *repositories.Repository
	auth *auth.JWTAuthenticator
}

type TokenData struct {
	AccessToken  string   `json:"access_token"`
	ExpiresIn    int      `json:"expires_in"`
	RefreshToken string   `json:"refresh_token"`
	Scope        []string `json:"scope"`
	TokenType    string   `json:"token_type"`
}

type UserData struct {
	ClientId  string   `json:"client_id"`
	Login     string   `json:"login"`
	Scopes    []string `json:"scopes"`
	UserId    string   `json:"user_id"`
	ExpiresIn int      `json:"expires_in"`
}

func NewTokenService(repo *repositories.Repository, auth *auth.JWTAuthenticator) *TokenService {
	return &TokenService{
		repo: repo,
		auth: auth,
	}
}

func (service *TokenService) ExchangeToken(code string) (*TokenData, error) {

	url := "https://id.twitch.tv/oauth2/token"

	body := map[string]string{
		"client_id":     os.Getenv("CLIENT_ID"),
		"client_secret": os.Getenv("CLIENT_SECRET"),
		"grant_type":    "authorization_code",
		"code":          code,
		"redirect_uri":  os.Getenv("REDIRECT_URI"),
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error encoding body to JSON: %v", err)
	}

	res, err := http.Post(url, "application/json", bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error sending POST request: %v", err)
	}
	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("twitch validation error: %v", string(respBody))
	}

	var tokenData TokenData
	if err := json.Unmarshal(respBody, &tokenData); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return &tokenData, nil
}

func (service *TokenService) ValidateToken(access_token string) (*UserData, error) {

	url := "https://id.twitch.tv/oauth2/validate"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error sending request to validate token: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("OAuth %v", access_token))

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to validate token: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("twitch validation error: %v", string(body))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var userData UserData
	if err := json.Unmarshal(body, &userData); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return &userData, nil
}

func (service *TokenService) SaveToken(userId, login, accessToken, refreshToken string) error {
	err := service.repo.Token.SaveToken(userId, login, accessToken, refreshToken)
	if err != nil {
		return err
	}
	return nil
}

func (service *TokenService) GenerateJWTToken(userId string) (string, error) {
	jwtToken, err := service.auth.GenerateToken(userId)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
