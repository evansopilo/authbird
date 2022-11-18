package main

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, resp map[string]interface{}) error {
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return err
	}
	return nil
}
