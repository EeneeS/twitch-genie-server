package services

import (
	"encoding/json"
	"fmt"
	"github.com/eenees/twitch-genie-server/internal/repositories"
	"io"
	"net/http"
	"os"
)

type ChannelService struct {
	repo *repositories.Repository
}

type ChannelData struct {
	Data       []Channel  `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Channel struct {
	Id    string `json:"broadcaster_id"`
	Login string `json:"broadcaster_login"`
	Name  string `json:"broadcaster_name"`
}

type Pagination struct {
	Cursors string `json:"cursor"`
}

func NewChannelService(repo *repositories.Repository) *ChannelService {
	return &ChannelService{repo: repo}
}

func (service *ChannelService) GetAccessToken(userId string) (string, error) {
	accessToken, err := service.repo.Token.GetAccessToken(userId)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (service *ChannelService) GetModeratedChannels(userId, accessToken string) (*ChannelData, error) {
	url := fmt.Sprintf("https://api.twitch.tv/helix/moderation/channels?user_id=%v", userId) // TODO: implement pagination or set max query to 100

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error sending GET request to get moderated channels: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", accessToken))
	req.Header.Set("Client-Id", os.Getenv("CLIENT_ID"))

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending GET request to get moderated channels: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("failed to get channels: %v", string(body))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var channelData ChannelData
	if err := json.Unmarshal(body, &channelData); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return &channelData, nil
}
