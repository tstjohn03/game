package models

import (
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	var conn, _ = upgrader.Upgrade(w, r, nil)
	balance := GetBalance()
	go func(conn *websocket.Conn) {
		for {
			ch := time.Tick(500 *time.Millisecond)

			for range ch {
				conn.WriteJSON(balance)
			}
		}
	}(conn)
}