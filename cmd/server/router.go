package main

import (
	"go-ticket/internal/handler"
	"net/http"
)

func SetupRouter(userHandler *handler.UserHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/profile", userHandler.GetProfile)
	mux.HandleFunc("/api/profile/update", userHandler.UpdateProfile)

	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	return mux
}
