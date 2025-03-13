package middleware 

import (
	"net/http"
	"strings"
	"log"
	"github.com/ganesh-kachare-josh/GameON/internal/pkg"
)
func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve token from the Authorization header

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: No token found", http.StatusUnauthorized)
			return
		}

		// Split to get the token part
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader { // If no "Bearer " prefix was found
			http.Error(w, "Unauthorized: Invalid token format", http.StatusUnauthorized)
			return
		}
		
		// Verify the token
		_, err := pkg.VerifyToken(token)
		if err != nil {
			log.Println(err)
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// Token is valid; proceed to the next handler
		next(w, r)
	}
}