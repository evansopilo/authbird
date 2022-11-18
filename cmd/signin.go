package main

import (
	"context"
	"encoding/json"
	"time"

	"log"
	"net/http"

	"github.com/evansopilo/authbird/pkg/data"
	"github.com/evansopilo/authbird/pkg/secure"
	"github.com/evansopilo/authbird/pkg/token"
	"github.com/golang-jwt/jwt"
)

func (app App) SignIn(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	w.Header().Set("Content-Type", "application/json")

	var authCred data.AuthCred

	if err := json.NewDecoder(r.Body).Decode(&authCred); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := map[string]interface{}{
			"status":  "error",
			"message": "provided invalid user data",
		}
		if err := WriteResponse(w, resp); err != nil {
			log.Printf("failed to create user, error: %v", err.Error())
			return
		}
	}

	user, err := app.Models.User.GetByEmail(ctx, authCred.Email)
	if err != nil || authCred.Email != user.Email {
		w.WriteHeader(http.StatusBadRequest)
		resp := map[string]interface{}{
			"status":  "error",
			"message": "prodived email already does exists",
		}
		if err := WriteResponse(w, resp); err != nil {
			log.Printf("failed to create user, error: %v", err.Error())
			return
		}
	}

	if err := secure.ComparePassword(authCred.Password, user.Password); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := map[string]interface{}{
			"status":  "error",
			"message": "prodived invalid user data",
		}
		if err := WriteResponse(w, resp); err != nil {
			log.Printf("failed to create user, error: %v", err.Error())
			return
		}
	}

	claims := token.Claims{UserID: user.ID, StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Second * 10).Unix()}}

	tokenStr, err := token.GenerateToken(claims, "secret")
	if err != nil {
		resp := map[string]interface{}{
			"status":  "error",
			"message": "prodived invalid user data",
		}
		if err := WriteResponse(w, resp); err != nil {
			log.Printf("failed to create user, error: %v", err.Error())
			return
		}
	}

	resp := map[string]interface{}{
		"status":  "success",
		"message": "user login successfull",
		"data": map[string]string{
			"token": *tokenStr,
		}}

	if err := WriteResponse(w, resp); err != nil {
		log.Printf("failed to authenticate user, error: %v", err.Error())
		return
	}
}
