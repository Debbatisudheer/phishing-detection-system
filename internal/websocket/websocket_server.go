package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Clients = make(map[*websocket.Conn]bool)

var Broadcast = make(chan []byte)

var upgrader = websocket.Upgrader{

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(
	w http.ResponseWriter,
	r *http.Request,
) {

	ws, err := upgrader.Upgrade(
		w,
		r,
		nil,
	)

	if err != nil {

		log.Println(err)

		return
	}

	defer ws.Close()

	Clients[ws] = true

	for {

		_, _, err := ws.ReadMessage()

		if err != nil {

			delete(Clients, ws)

			break
		}
	}
}

func HandleMessages() {

	for {

		msg := <-Broadcast

		for client := range Clients {

			err := client.WriteMessage(
				websocket.TextMessage,
				msg,
			)

			if err != nil {

				client.Close()

				delete(Clients, client)
			}
		}
	}
}