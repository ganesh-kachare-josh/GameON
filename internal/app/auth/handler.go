package auth

import (
    "context"
    "encoding/json"
    "net/http"
    "strings"
    "time"

    "github.com/ganesh-kachare-josh/GameON/internal/pkg"
    "github.com/ganesh-kachare-josh/GameON/internal/repository"
)

func Login(authService Service) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {

        var login repository.Login

        err := json.NewDecoder(r.Body).Decode(&login)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        ctx := context.Background()

        response, err := authService.Login(ctx, login)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        cookie := http.Cookie{
            Name:     "token",
            Value:    response.Token,
            Expires:  time.Now().Add(24 * time.Hour),
            HttpOnly: true,
            Secure:   false,
            Path:     "/",
            SameSite: http.SameSiteLaxMode,
        }

        http.SetCookie(w, &cookie)

        w.WriteHeader(http.StatusOK)
        err = json.NewEncoder(w).Encode(response)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}

func Register(authService Service) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {

        var register RegisterData

        err := json.NewDecoder(r.Body).Decode(&register)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        ctx := context.Background()

        register, err = authService.Register(ctx, register)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)

        err = json.NewEncoder(w).Encode(register)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

    }
}

func Logout(authService Service) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {

        ctx := r.Context()
        response := authService.Logout(ctx)

        cookie := http.Cookie{
            Name:     "auth_token",
            Value:    "",
            HttpOnly: true,
            Secure:   false,
            Path:     "/",
            SameSite: http.SameSiteLaxMode,
            MaxAge:   -1,
        }

        http.SetCookie(w, &cookie)

        w.WriteHeader(http.StatusOK)
        err := json.NewEncoder(w).Encode(response)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

    }
}

func IsLogin(authService Service) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {

        var loginstatus LoginStatus

        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            loginstatus.User_id = 0
            loginstatus.Islogin = false
            err := json.NewEncoder(w).Encode(loginstatus)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
            return
        }

        // Split to get the token part
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader { // If no "Bearer " prefix was found
            loginstatus.User_id = 0
            loginstatus.Islogin = false
            err := json.NewEncoder(w).Encode(loginstatus)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
            return
        }
		
        user_id, err := pkg.GetUserIdFromToken(tokenString)

        if err != nil {
            loginstatus.User_id = 0
            loginstatus.Islogin = false
            err := json.NewEncoder(w).Encode(loginstatus)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
            return
        }

        loginstatus.User_id = user_id
        loginstatus.Islogin = true
        err = json.NewEncoder(w).Encode(loginstatus)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

    }
}