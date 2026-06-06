package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Clients = make(
	map[*websocket.Conn]bool,
)

var Broadcast =
	make(chan []byte)

var upgrader =
	websocket.Upgrader{

		CheckOrigin: func(
			r *http.Request,
		) bool {

			return true
		},
	}

func HandleConnections(
	w http.ResponseWriter,
	r *http.Request,
) {

	log.Println(
		"New WebSocket Connection",
	)

	ws, err :=
		upgrader.Upgrade(
			w,
			r,
			nil,
		)

	if err != nil {

		log.Println(
			"Upgrade Error:",
			err,
		)

		return
	}

	Clients[ws] = true

	log.Println(
		"Client Connected",
	)

	defer func() {

		delete(
			Clients,
			ws,
		)

		ws.Close()

		log.Println(
			"Client Disconnected",
		)
	}()

	for {

		_, _, err :=
			ws.ReadMessage()

		if err != nil {

			log.Println(
				"Read Error:",
				err,
			)

			break
		}
	}
}

func HandleMessages() {

	for {

		msg :=
			<-Broadcast

		log.Println(
			"Broadcasting:",
			string(msg),
		)

		for client := range Clients {

			err :=
				client.WriteMessage(
					websocket.TextMessage,
					msg,
				)

			if err != nil {

				log.Println(
					"Write Error:",
					err,
				)

				client.Close()

				delete(
					Clients,
					client,
				)
			}
		}
	}
}