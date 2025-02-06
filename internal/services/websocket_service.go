package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/eenees/twitch-genie-server/internal/repositories"
	"github.com/gorilla/websocket"
)

type WebsocketService struct {
  repo *repositories.Repository
}

func NewWebSocketService(repo *repositories.Repository) *WebsocketService {
  return &WebsocketService{repo: repo}
}

func (service *WebsocketService) IsChannelModerator(channelId, userId string) (bool, error) {
  url := fmt.Sprintf("https://api.twitch.tv/helix/moderation/channels?user_id=%v", userId)

  client := &http.Client{}

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return false, fmt.Errorf("error sending request to validate if moderator: %v", err)
  }

  accessToken, err := service.repo.Token.GetAccessToken(userId)
  if err != nil {
    return false, fmt.Errorf("error getting accessToken: %v", err)
  }

  clientId := os.Getenv("CLIENT_ID")

  req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", accessToken))
  req.Header.Set("Client-Id", clientId)

  res, err := client.Do(req)
  if err != nil {
    return false, fmt.Errorf("error sending request to validate if moderator: %v", err)
  }
  defer res.Body.Close()

  if res.StatusCode != http.StatusOK {
    body, _ := io.ReadAll(res.Body)
    return false, fmt.Errorf("twitch error: %v", body)
  }

  body, err := io.ReadAll(res.Body)
  if err != nil {
    fmt.Printf("error reading response body: %v", err)
    return false, fmt.Errorf("error reading response body: %v", err)
  }

  var channelData ChannelData
  if err := json.Unmarshal(body, &channelData); err != nil {
    return false, fmt.Errorf("error decoding response body: %v", err)
  }

  for _, channel := range channelData.Data {
    if channel.Id == channelId {
      return true, nil
    }
  }

  return false, fmt.Errorf("you don't moderate the given channel id")
}

type BaseMessage struct {
  Type string `json:"type"`
}

type ImageMessage struct {
  BaseMessage
  Xpos  int `json:"y_pos"`
  Ypos int `json:"x_pos"`
  Event string `json:"event"`
}

type SoundMessage struct {
  BaseMessage
  Source string `json:"source"`
}

func (service *WebsocketService) ReadMessage(conn *websocket.Conn) (interface{}, error) {
  _, msg, err := conn.ReadMessage()
  if err != nil {
    if closeErr, ok := err.(*websocket.CloseError); ok && closeErr.Code == websocket.CloseNormalClosure {
      return nil, fmt.Errorf("non closing error")
    }
    fmt.Println("Read error:", err)
    return nil, fmt.Errorf("read error")
  }

  var baseMessage BaseMessage
  if err := json.Unmarshal(msg, &baseMessage); err != nil {
    return nil, fmt.Errorf("unmarshal error")
  }

  var processedMessage interface{}
  switch baseMessage.Type {
  case "image":
    var m ImageMessage
    if err := json.Unmarshal(msg, &m); err != nil {
      return nil, fmt.Errorf("unmarshal error")
    }
    processedMessage = m
  default:
    return nil, fmt.Errorf("invalid message type")
  }

  return processedMessage, nil
}
