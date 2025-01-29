package main

import "net/http"

func (app *application) exchangeTokenHandler(w http.ResponseWriter, r *http.Request) {
	// 1. retrieve the `code` from the body.
	// 2. do a post request to: https://id.twitch.tv/oauth2/token with this query.
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
