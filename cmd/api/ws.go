package main

import (
  "fmt"
  "net/http"
  "sync"

  "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
  CheckOrigin: func(r *http.Request) bool {
    return true // FIX: change
  },
}

var connections = make(map[string][]*websocket.Conn)
var mu sync.Mutex

// TODO: validate the data being send?
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

  mu.Lock()
  connections[channelId] = append(connections[channelId], conn)
  mu.Unlock()

  for {
    _, msg, err := conn.ReadMessage()
    if err != nil {
      if closeErr, ok := err.(*websocket.CloseError); ok && closeErr.Code == websocket.CloseNormalClosure {
        break
      }
      fmt.Println("Read error:", err)
      break
    }

    mu.Lock()
    for _, c := range connections[channelId] {
      if err := c.WriteMessage(websocket.TextMessage, msg); err != nil {
        fmt.Println("Write error:", err)
        break
      }
    }
    mu.Unlock()
  }

	mu.Lock()
	defer func() {
		for i, c := range connections[channelId] {
			if c == conn {
				connections[channelId] = append(connections[channelId][:i], connections[channelId][i+1:]...)
				break
			}
		}
		mu.Unlock()
	}()
}
