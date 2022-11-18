package main

import (
	"log"
	"net/http"
)

func (app App) Health(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	resp := map[string]interface{}{
		"status":  "success",
		"message": "health",
	}

	if err := WriteResponse(w, resp); err != nil {
		log.Printf("failed to get health info, error: %v", err.Error())
		return
	}
}
