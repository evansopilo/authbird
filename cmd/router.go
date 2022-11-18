package main

import (
	"github.com/gorilla/mux"
)

func (app App) Router() *mux.Router {
	mux := mux.NewRouter()

	mux.HandleFunc("/api/health", app.Health)

	mux.HandleFunc("/api/signup", app.SignUp)

	mux.HandleFunc("/api/signin", app.SignIn)

	mux.HandleFunc("/api/verify", app.VerifyToken)

	// Wildcard route /api/*
	mux.HandleFunc("/api/", app.Wildcard)

	return mux
}
