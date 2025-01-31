package middlewares

import (
	"context"
	"github.com/eenees/twitch-genie-server/internal/utils/auth"
	"net/http"
)

func AuthMiddleware(auth *auth.JWTAuthenticator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("token")
			if err != nil {
				http.Error(w, "Unauthorized: Cookie not found", http.StatusUnauthorized)
				return
			}

			jwtToken, err := auth.VerifyToken(cookie.Value)
			if err != nil {
				http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
				return
			}

			userId, err := jwtToken.Claims.GetSubject()
			if err != nil {
				http.Error(w, "Unauthorized: No ID found in token", http.StatusUnauthorized)
			}

			ctx := context.WithValue(r.Context(), "token", userId)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

