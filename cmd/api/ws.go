package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // FIX: change
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
}
