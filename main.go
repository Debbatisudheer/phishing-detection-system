package main

import (
	"fmt"
	"net/http"

	"phishing-platform/routes"
	"phishing-platform/database"
	"phishing-platform/internal/smtpserver"
	"github.com/rs/cors"
	"phishing-platform/internal/websocket"
)

func main() {

	database.ConnectDatabase()

	go smtpserver.StartSMTPServer()
	go websocket.HandleMessages()

	routes.SetupRoutes()

	fmt.Println("Server running on port 8081")

	handler := cors.Default().Handler(http.DefaultServeMux)

http.ListenAndServe(
	":8081",
	handler,
)

http.HandleFunc(
	"/ws",
	websocket.HandleConnections,
)
}