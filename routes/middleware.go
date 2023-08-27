package routes

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, helloerror := r.Cookie("token")
		// if tokenString == "" {
		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// 	return
		// }
		if helloerror != nil {
			if helloerror == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tokenStr := tokenString.Value
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Token is valid, call the next handler
		next.ServeHTTP(w, r)
	})
}
