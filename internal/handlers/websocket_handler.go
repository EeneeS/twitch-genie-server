package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/eenees/twitch-genie-server/internal/services"
	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
  service *services.WebsocketService
}

func NewWebSocketHandler(service *services.WebsocketService) *WebsocketHandler {
  return &WebsocketHandler{service: service}
}

var upgrader = websocket.Upgrader{
  CheckOrigin: func(r *http.Request) bool {
    return true // FIX: change
  },
}

var connections = make(map[string][]*websocket.Conn)
var mu sync.Mutex

func (handler *WebsocketHandler) Init(w http.ResponseWriter, r *http.Request) {
  userId, ok := r.Context().Value("userId").(string)
  if !ok {
    http.Error(w, "token not found in request", http.StatusBadRequest)
    return
  }

  channelId := r.URL.Query().Get("channel_id")
  if channelId == "" {
    http.Error(w, "missing channel id", http.StatusBadRequest)
    return
  }

  _, err := handler.service.IsChannelModerator(channelId, userId)
  if err != nil {
    http.Error(w, "unauthorized", http.StatusUnauthorized)
    return
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

  /*
  type: 'image' | 'sound"
  */

  for {

    message, err := handler.service.ReadMessage(conn)
    if err != nil {
      fmt.Println(err.Error())
    }
    fmt.Println(message)

    // mu.Lock()
    // for _, c := range connections[channelId] {
    //   if err := c.WriteMessage(websocket.TextMessage, msg); err != nil {
    //     fmt.Println("Write error:", err)
    //     break
    //   }
    // }
    // mu.Unlock()
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
