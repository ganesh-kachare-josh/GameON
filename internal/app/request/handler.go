package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ganesh-kachare-josh/GameON/internal/pkg"

)

func GetRequestById(requestService Service)(func (w http.ResponseWriter , r *http.Request)) {
	return func(w http.ResponseWriter , r *http.Request) {
		ctx := r.Context() 

		vars := mux.Vars(r)
		id := vars["id"]
		if id == "" {
			http.Error(w,"id is required",http.StatusBadRequest)
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
			http.Error(w,"id is required",http.StatusBadRequest)
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
			http.Error(w,"id is required",http.StatusBadRequest)
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
			http.Error(w ,"failed to decode request body",http.StatusInternalServerError) 
			return 
		}

		body.Request_id = request_id 

		response , emailResponse , err := acceptRequestService.AcceptRequest(ctx , body )
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

		err = pkg.SendEmail(emailResponse.Email, 
		fmt.Sprintf("🎉 Game On! %v Your Request Was Accepted!" , emailResponse.CreatorName), 
		fmt.Sprintf("Great news! %v has accepted your game request to play %v. Get ready to jump into action and enjoy the thrill! 🚀\n\nLog in now to check the details and start gaming!\n\nHappy Gaming! 🎮" , emailResponse.ParticipantName , emailResponse.Sport),
		)
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
			http.Error(w,"id is required",http.StatusBadRequest)
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
			http.Error(w ,"failed to decode request body",http.StatusInternalServerError) 
			return 
		}

		body.Request_id = request_id 

		response , emailResponse , err  := confirmRequestService.ConfirmRequest(ctx , body) 
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

		err = pkg.SendEmail(emailResponse.Email, 
		"✅ You're In! Join Request Confirmed", 
		fmt.Sprintf("Congratulations! %v has accepted your join request to play %v. You're now part of the squad! 🔥\n\nPrepare yourself, gear up, and get ready for an epic gaming session.\n\nSee you in the game! 🎮",emailResponse.CreatorName , emailResponse.Sport),
		)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
	}
}

func DeleteRequest(deleteRequest Service) (func (w http.ResponseWriter , r *http.Request)) {
	return func(w http.ResponseWriter , r * http.Request) {
		ctx := r.Context() 

		vars := mux.Vars(r)
		id := vars["request_id"]
		if id == "" {
			http.Error(w,"id is required",http.StatusBadRequest)
			return 
		}

		request_id,err := strconv.Atoi(id)
		if err != nil {
			http.Error(w,err.Error(),http.StatusBadRequest)
			return
		}

		result , err := deleteRequest.DeleteRequest(ctx , request_id)
		if err != nil {
			http.Error(w,fmt.Sprintf("failed to delete request: %v", err),http.StatusInternalServerError)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			http.Error(w, fmt.Sprintf("error checking rows affected: %v", err), http.StatusInternalServerError)
			return
		}

		if rowsAffected == 0 {
			http.Error(w, "Request not found", http.StatusNotFound)
			return
		}
		msg := DeleteResponse{
			Message:  "Play request deleted successfully",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(msg)

	}
}

func RejectParticipant(rejectRequest Service) (func (w http.ResponseWriter , r *http.Request)) {
	return func(w http.ResponseWriter , r * http.Request) {
		ctx := r.Context() 

		vars := mux.Vars(r)
		id := vars["pid"]
		if id == "" {
			http.Error(w,"id is required",http.StatusBadRequest)
			return 
		}
		participant_id,err := strconv.Atoi(id)
		if err != nil {
			http.Error(w,err.Error(),http.StatusBadRequest)
			return
		} 

		emailResponse , err := rejectRequest.RejectParticipant(ctx , participant_id)
		if err != nil {
			http.Error(w,fmt.Sprintf("failed to delete request: %v", err),http.StatusInternalServerError)
			return
		}

		
		msg := DeleteResponse{
			Message:  "Participant has been rejected from the request.",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(msg)

		err = pkg.SendEmail(emailResponse.Email, 
			fmt.Sprintf("❌ Oops! Join Request Rejected by %v" , emailResponse.CreatorName), 
			fmt.Sprintf("Hey there, unfortunately, your request to join the game %v was not accepted this time. But don’t worry, new opportunities are always around the corner! 🌟\n\nKeep exploring, find another game, and show them what they’re missing!\n\nBetter luck next time! 🎮",emailResponse.Sport), 
		)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
		
	}
}

func CreateRequest(createRequest Service) func (w http.ResponseWriter , r *http.Request) {
	return func (w http.ResponseWriter , r *http.Request) {
		ctx := r.Context()

		var requestBody Request 

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w , fmt.Sprintf("failed to decode the request body: %v" , err) , http.StatusInternalServerError)
			return
		}

		response , err := createRequest.CreateRequest(ctx , requestBody)
		if err != nil {
			http.Error(w , fmt.Sprintf("failed to create request: %v" ,err ) , http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {	
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}

	}
}