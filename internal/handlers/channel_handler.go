package handlers

import (
	"encoding/json"
	"github.com/eenees/twitch-genie-server/internal/services"
	"net/http"
)

type ChannelHandler struct {
	service *services.ChannelService
}

func NewChannelHandler(service *services.ChannelService) *ChannelHandler {
	return &ChannelHandler{service: service}
}

// GetModeratedChanels godoc
//
// @Summary get channels you moderate
// @Description get information about the channels you moderate
// @Produce json
// @router /moderated-channels [get]
// @Security ApiKeyAuth
// @Tags Channels
func (handler *ChannelHandler) GetModeratedChannels(w http.ResponseWriter, r *http.Request) {

	userId, ok := r.Context().Value("userId").(string)
	if !ok {
		http.Error(w, "token not found in request", http.StatusUnauthorized)
		return
	}

	accessToken, err := handler.service.GetAccessToken(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	channels, err := handler.service.GetModeratedChannels(userId, accessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(channels)
}
