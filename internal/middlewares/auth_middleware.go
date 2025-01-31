package middlewares

import (
	"context"
	"net/http"

	"github.com/eenees/twitch-genie-server/internal/utils/auth"
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

			ctx := context.WithValue(r.Context(), "token", jwtToken)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

