package services

import "github.com/eenees/twitch-genie-server/internal/repositories"

type ChannelService struct {
	repo *repositories.Repository
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
