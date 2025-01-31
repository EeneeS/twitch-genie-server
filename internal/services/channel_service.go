package services

import "github.com/eenees/twitch-genie-server/internal/repositories"

type ChannelService struct {
	repo *repositories.Repository
}

func NewChannelService(repo *repositories.Repository) *ChannelService {
	return &ChannelService{repo: repo}
}
