package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	specs "github.com/purisaurabh/database-connection/internal/pkg"
)

func decodePostRequest(r *http.Request) (specs.PostProfile, error) {
	var req specs.PostProfile
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("Error decoding request:", err)
		return specs.PostProfile{}, err
	}
	return req, nil
}

func decodeUpdateRequest(r *http.Request) (specs.UpdateProfile, error) {
	var req specs.UpdateProfile
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("Error decoding request:", err)
		return specs.UpdateProfile{}, err
	}
	return req, nil
}

func GetIDFromRequest(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return 0, fmt.Errorf("profile_id not found")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, fmt.Errorf("profile_id not found")
	}
	return idInt, nil
}
