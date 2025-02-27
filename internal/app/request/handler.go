package request

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func GetRequestById(requestService Service)(func (w http.ResponseWriter , r *http.Request)) {
	return func(w http.ResponseWriter , r *http.Request) {
		ctx := r.Context() 

		vars := mux.Vars(r)
		id := vars["id"]
		if id == "" {
			http.Error(w,errors.New("id is required").Error(),http.StatusBadRequest)
			return 
		}
		request_id,err := strconv.Atoi(id)
		if err != nil {
			http.Error(w,err.Error(),http.StatusBadRequest)
			return
		} 

		response , err := requestService.GetRequestById(ctx , request_id) 
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return 
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) 
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}		
	}
}

func GetAllRequests(requestService Service)(func (w http.ResponseWriter , r *http.Request)) {
	return func(w http.ResponseWriter , r *http.Request) {
		ctx := r.Context() 

		response , err := requestService.GetAllRequests(ctx) 
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return 
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) 
		
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}		
	}
}

func GetAllParticipants(participantService Service) (func (w http.ResponseWriter , r *http.Request)) {
	return func(w http.ResponseWriter , r *http.Request) {
		ctx := r.Context() 

		vars := mux.Vars(r)
		id := vars["id"]
		if id == "" {
			http.Error(w,errors.New("id is required").Error(),http.StatusBadRequest)
			return 
		}

		request_id,err := strconv.Atoi(id)
		if err != nil {
			http.Error(w,err.Error(),http.StatusBadRequest)
			return
		} 
		
		response := participantService.GetAllParticipants(ctx , request_id)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) 
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}

	}
}

func AcceptRequest(acceptRequestService Service) (func (w http.ResponseWriter , r * http.Request)) {
	return func (w http.ResponseWriter , r *http.Request) {
		ctx := r.Context()

		vars := mux.Vars(r)
		id := vars["request_id"]
		if id == "" {
			http.Error(w,errors.New("id is required").Error(),http.StatusBadRequest)
			return 
		}

		request_id,err := strconv.Atoi(id)
		if err != nil {
			http.Error(w,err.Error(),http.StatusBadRequest)
			return
		}

		var body AcceptRequestBody 
		err = json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w , errors.New("failed to decode request body").Error(),http.StatusInternalServerError) 
			return 
		}

		body.Request_id = request_id 

		response := acceptRequestService.AcceptRequest(ctx , body ) 

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) 
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
	}
}

func ConfirmRequest(confirmRequestService Service) (func (w http.ResponseWriter , r * http.Request)) {
	return func (w http.ResponseWriter , r *http.Request) {
		ctx := r.Context()

		vars := mux.Vars(r)
		id := vars["request_id"]
		if id == "" {
			http.Error(w,errors.New("id is required").Error(),http.StatusBadRequest)
			return 
		}

		request_id,err := strconv.Atoi(id)
		if err != nil {
			http.Error(w,err.Error(),http.StatusBadRequest)
			return
		}

		var body AcceptRequestBody 
		err = json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w , errors.New("failed to decode request body").Error(),http.StatusInternalServerError) 
			return 
		}

		body.Request_id = request_id 

		response := confirmRequestService.ConfirmRequest(ctx , body) 

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) 
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
	}
}