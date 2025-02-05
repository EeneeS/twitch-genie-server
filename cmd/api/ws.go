package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
  CheckOrigin: func(r *http.Request) bool {
    return true // FIX: change
  },
}

// WebSocket handler
func (app *application) wsHandler(w http.ResponseWriter, r *http.Request) {

  channelId := r.URL.Query().Get("channel_id")
  if channelId == "" {
    http.Error(w, "missing channel id", http.StatusBadRequest)
  }

  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  defer conn.Close()

  for {
    _, msg, err := conn.ReadMessage()
    if err != nil {
      fmt.Println("Read error:", err)
      break
    }

    fmt.Printf("Received: %s\n", msg)

    if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
      fmt.Println("Write error:", err)
      break
    }
  }
}
