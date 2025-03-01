package middleware 

import "net/http"

// CORS allows cross-origin requests from any origin (i.e., *).
func CORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Allow all origins
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle pre-flight OPTIONS request
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        // Continue with the next handler
        next.ServeHTTP(w, r)
    })
}