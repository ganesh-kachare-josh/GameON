package profile

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"github.com/ganesh-kachare-josh/GameON/internal/pkg"

)

func GetUserById(profileService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		vars := mux.Vars(r)
		id := vars["user_id"]
		if id == "" {
			http.Error(w, "id is required", http.StatusBadRequest)
			return
		}
		user_id, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := profileService.GetUserById(ctx, user_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func UpdateProfile(profileService Service) func(w http.ResponseWriter , r *http.Request) {
	return func(w http.ResponseWriter , r *http.Request) {
		ctx := r.Context() 

		authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "invalid authorization header", http.StatusInternalServerError)   
            return
		}

        // Split to get the token part
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader { // If no "Bearer " prefix was found
            http.Error(w, "invalid header formate", http.StatusInternalServerError)
            return
        }
		
        user_id, err := pkg.GetUserIdFromToken(tokenString)
		if err != nil {
			http.Error(w , err.Error() , http.StatusInternalServerError)
			return
		}

		var requestBody UserData 
		err = json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w , "error decoding the request body" , http.StatusInternalServerError)
			return 
		}

		requestBody.Id = user_id 


		response , err := profileService.UpdateProfile(ctx , requestBody)
		if err != nil {
			http.Error(w , err.Error() , http.StatusInternalServerError)
			return 
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w , "error in encoding the response" , http.StatusInternalServerError)
			return 
		}

	}
}
