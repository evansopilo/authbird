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

func (app App) SignUp(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	w.Header().Set("Content-Type", "application/json")

	var user data.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := map[string]interface{}{
			"status":  "error",
			"message": "status bad request",
		}
		if err := WriteResponse(w, resp); err != nil {
			log.Printf("failed to create user, error: %v", err.Error())
			return
		}
		log.Printf("failed to create user, error: %v", err.Error())
		return
	}

	existingUser, err := app.Models.User.GetByEmail(ctx, user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := map[string]interface{}{
			"status":  "error",
			"message": "failed to create user",
		}
		if err := WriteResponse(w, resp); err != nil {
			log.Printf("failed to create user, error: %v", err.Error())
			return
		}
	}

	if existingUser.Email == user.Email {
		w.WriteHeader(http.StatusBadRequest)
		resp := map[string]interface{}{
			"status":  "error",
			"message": "provided email address already exists",
		}
		if err := WriteResponse(w, resp); err != nil {
			log.Printf("failed to create user, error: %v", err.Error())
			return
		}
	}

	hashedPassword, err := secure.HashPassword(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := map[string]interface{}{
			"status":  "error",
			"message": "failed to create user",
		}
		if err := WriteResponse(w, resp); err != nil {
			log.Printf("failed to hash user provided password, error: %v", err.Error())
			return
		}
		log.Printf("failed to hash user provided password, error: %v", err.Error())
		return
	}

	user.Password = *hashedPassword

	if err := app.Models.User.Create(ctx, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := map[string]interface{}{
			"status":  "error",
			"message": "failed to create user",
		}
		if err := WriteResponse(w, resp); err != nil {
			log.Printf("failed to hash user provided password, error: %v", err.Error())
			return
		}
		log.Printf("failed to hash user provided password, error: %v", err.Error())
		return
	}

	claims := token.Claims{UserID: user.ID, StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Second * 10).Unix()}}

	token, err := token.GenerateToken(claims, "secret")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := map[string]interface{}{
			"status":  "error",
			"message": "failed to create user",
		}
		if err := WriteResponse(w, resp); err != nil {
			log.Printf("failed to hash user provided password, error: %v", err.Error())
			return
		}
		log.Printf("failed to hash user provided password, error: %v", err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)

	resp := map[string]interface{}{
		"status":  "success",
		"message": "user created successfull",
		"data": map[string]string{
			"token": *token,
		},
	}
	if err := WriteResponse(w, resp); err != nil {
		log.Printf("failed to create user, error: %v", err.Error())
		return
	}
}
