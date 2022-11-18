package main

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/evansopilo/authbird/pkg/data"
	"github.com/evansopilo/authbird/pkg/token"
)

func (app App) VerifyToken(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var tokenStr data.TokenStr

	if err := json.NewDecoder(r.Body).Decode(&tokenStr); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := map[string]interface{}{
			"status":  "error",
			"message": "bad service request request",
		}
		if err := WriteResponse(w, resp); err != nil {
			log.Printf("failed to verify provided token string, error: %v", err.Error())
		}
		return
	}

	claims, err := token.Verify(tokenStr.Value, "secret")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		resp := map[string]interface{}{
			"status":  "error",
			"message": "token verification error",
		}
		if err := WriteResponse(w, resp); err != nil {
			log.Printf("failed to verify provided token string, error: %v", err.Error())
		}
		return
	}

	resp := map[string]interface{}{
		"status":  "success",
		"message": "token verification successfull",
		"data": map[string]interface{}{
			"claims": claims,
		}}

	if err := WriteResponse(w, resp); err != nil {
		log.Printf("failed to verify provided token string, error: %v", err.Error())
	}
}
