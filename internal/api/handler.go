package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/purisaurabh/database-connection/internal/app"
)

func PostDataApi(ctx context.Context, service app.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Decode request
		req, err := decodePostRequest(r)
		if err != nil {
			fmt.Println("Error decoding request:", err)
			ErrorResponse(w, http.StatusBadRequest, Response{Message: "Error decoding request"})
			return
		}

		// Validate request
		err = req.Validate()
		if err != nil {
			fmt.Println("Error validating request:", err)
			ErrorResponse(w, http.StatusBadRequest, Response{Message: "Error validating request"})
			return
		}

		// Call service
		err = service.PostProfileData(ctx, req)
		if err != nil {
			fmt.Println("Error calling service:", err)
			ErrorResponse(w, http.StatusInternalServerError, Response{Message: "Error posting data"})
			return
		}

		// Encode response
		SuccessResponse(w, http.StatusOK, Response{Message: "Data posted successfully"})
	}
}

func ListProfileData(ctx context.Context, service app.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Call service
		resp, err := service.ListProfileData(ctx)
		if err != nil {
			fmt.Println("Error calling service:", err)
			ErrorResponse(w, http.StatusInternalServerError, Response{Message: "Error listing data"})
			return
		}

		// Encode response
		SuccessResponse(w, http.StatusOK, resp)
	}
}

func UpdateProfileData(ctx context.Context, service app.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := GetIDFromRequest(r)
		if err != nil {
			fmt.Println("Error getting id from request:", err)
			ErrorResponse(w, http.StatusBadRequest, Response{Message: "Error getting id from request"})
			return
		}
		// Decode request
		req, err := decodeUpdateRequest(r)
		if err != nil {
			fmt.Println("Error decoding request:", err)
			ErrorResponse(w, http.StatusBadRequest, Response{Message: "Error decoding request"})
			return
		}

		// Validate request
		err = req.Validate()
		if err != nil {
			fmt.Println("Error validating request:", err)
			ErrorResponse(w, http.StatusBadRequest, Response{Message: "Error validating request"})
			return
		}

		// Call service
		err = service.UpdateProfileData(ctx, id, req)
		if err != nil {
			fmt.Println("Error calling service:", err)
			ErrorResponse(w, http.StatusInternalServerError, Response{Message: "Error updating data"})
			return
		}

		// Encode response
		SuccessResponse(w, http.StatusOK, Response{Message: "Data updated successfully"})
	}
}
