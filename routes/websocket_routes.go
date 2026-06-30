package routes

import (
	"net/http"

	"phishing-platform/internal/websocket"
)

func RegisterWebsocketRoutes() {

	http.HandleFunc(
		"/ws",
		websocket.HandleConnections,
	)

}