package main

import (
	"fmt"
	"net/http"

	"phishing-platform/routes"
	"phishing-platform/database"
	"phishing-platform/internal/smtpserver"
	"phishing-platform/internal/websocket"
	"phishing-platform/internal/sandbox"

	"github.com/rs/cors"
)

func main() {

	database.ConnectDatabase()

	sandbox.StartSandboxWorker()

	go smtpserver.StartSMTPServer()

	go websocket.HandleMessages()

	routes.SetupRoutes()

	fmt.Println(
		"Server running on port 8081",
	)

	c := cors.New(
		cors.Options{
			AllowedOrigins: []string{
				"http://localhost:5173",
			},
			AllowedMethods: []string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"OPTIONS",
			},
			AllowedHeaders: []string{
				"*",
			},
			AllowCredentials: true,
		},
	)

	handler := c.Handler(
		http.DefaultServeMux,
	)

	http.ListenAndServe(
		":8081",
		handler,
	)
}