package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ExchangeTokenRequestBody struct {
	Code string `json:"code" binding:"required"`
}

type UserData struct {
	Login string `json:"login"`
}

type ExchangeTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// ExchangeToken godoc
//
// @Summary Exchange token
// @Description Exchange the auth token and retrieve user data
// @Accepts json
// @Produce json
// @Param exchangeTokenBody body ExchangeTokenRequestBody true "Exchange token request"
// @router /exchange-token [post]
func (app *application) exchangeTokenHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var body ExchangeTokenRequestBody

	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}

	if body.Code == "" {
		http.Error(w, "code is required", http.StatusBadRequest)
		return
	}

	tokenResponse, err := exchangeToken(body.Code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(tokenResponse.AccessToken) // zodatn stopt me zagen
	return

	//FIX: veel te veel werk gehad aan deze stappen uitsvhrijvefndskfdf

	// 1. retrieve the `code` from the body. V
	// 2. do a post request to: https://id.twitch.tv/oauth2/token with this query. V
	// client_id=hof5gwx0su6owfnys0yan9c87zr6t
	// &client_secret=41vpdji4e9gif29md0ouet6fktd2
	// &code=gulfwdmys5lsm6qyz4xiz9q32l10
	// &grant_type=authorization_code
	// &redirect_uri=http://localhost:3000
	// 3. read the response from your post request
	//   {
	//   "access_token": "rfx2uswqe8l4g1mkagrvg5tv0ks3",
	//   "expires_in": 14124,
	//   "refresh_token": "5b93chm6hdve3mycz05zfzatkfdenfspp1h1ar2xxdalen01",
	//   "scope": [
	//     "channel:moderate",
	//     "chat:edit",
	//     "chat:read"
	//   ],
	//   "token_type": "bearer"
	// }
	// and extract the access_token, refresh_type
	// 4. validate the user with get request to: https://id.twitch.tv/oauth2/validate
	//   curl -X GET 'https://id.twitch.tv/oauth2/validate' \
	// -H 'Authorization: OAuth <access token to validate goes here>'
	// 5. read the response from the validate request.
	//   {
	//   "client_id": "wbmytr93xzw8zbg0p1izqyzzc5mbiz",
	//   "login": "twitchdev",
	//   "scopes": [
	//     "channel:read:subscriptions"
	//   ],
	//   "user_id": "141981764",
	//   "expires_in": 5520838
	// }
	// 6. get the login and user_id from the request.
	// 7. save { user_id, login, access_token, refresh_token } to database.
	// 8. return the login and/or user_id to the user.
}

func exchangeToken(code string) (*ExchangeTokenResponse, error) {
	baseUrl := os.Getenv("EXCHANGE_TOKEN_URL")
	body := map[string]string{
		"client_id":     os.Getenv("CLIENT_ID"),
		"client_secret": os.Getenv("CLIENT_SECRET"),
		"grant_type":    "authorization_code",
		"code":          code,
		"redirect_uri":  os.Getenv("REDIRECT_URI"),
	}

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	err := encoder.Encode(body)
	if err != nil {
		return nil, fmt.Errorf("error encoding body to JSON: %v", err)
	}

	response, err := http.Post(baseUrl, "application/json", &buf)
	if err != nil {
		return nil, fmt.Errorf("error sending POST request: %v", err)
	}
	defer response.Body.Close()

	bodyResp, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		var errorBody struct {
			Message string `json:message`
			Status  int    `json:"status"`
		}

		if err := json.Unmarshal(bodyResp, &errorBody); err != nil {
			return nil, fmt.Errorf("twitch responded with status %v but failed to parse error body: %v", response.StatusCode, err)
		}

		return nil, fmt.Errorf("twitch error: %s", errorBody.Message)
	}

	var tokenResponse ExchangeTokenResponse
	if err := json.Unmarshal(bodyResp, &tokenResponse); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return &tokenResponse, nil
}
