package main

import (
	"fmt"
	"log"
	"net/http"

	"phishing-platform/database"
	"phishing-platform/internal/sandbox"
	"phishing-platform/internal/smtpserver"
	"phishing-platform/internal/websocket"
	"phishing-platform/routes"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load()

	if err != nil {

		log.Println(
			".env file not loaded",
		)

	} else {

		fmt.Println(
			".env loaded successfully",
		)
	}

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