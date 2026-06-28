package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	// -------------------------
	// Database
	// -------------------------

	database.ConnectDatabase()

	// -------------------------
	// Background Workers
	// -------------------------

	sandbox.StartSandboxWorker()

	database.StartCleanupWorker()

	go smtpserver.StartSMTPServer()

	go websocket.HandleMessages()

	// -------------------------
	// Routes
	// -------------------------

	routes.SetupRoutes()

	// -------------------------
	// Railway / Local Port
	// -------------------------

	port := os.Getenv("PORT")

	if port == "" {

		port = "8081"

	}

	fmt.Println(
		"Server running on port",
		port,
	)

	// -------------------------
	// CORS
	// -------------------------

	c := cors.New(

		cors.Options{

			AllowedOrigins: []string{

				"http://localhost:5173",

				// Add your Vercel URL here later.
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

	// -------------------------
	// Start Server
	// -------------------------

	log.Fatal(

		http.ListenAndServe(

			":"+port,

			handler,
		),
	)
}