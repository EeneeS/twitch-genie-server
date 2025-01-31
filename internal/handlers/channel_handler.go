package handlers

import (
	"fmt"
	"net/http"

	"github.com/eenees/twitch-genie-server/internal/services"
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

	token, ok := r.Context().Value("token").(string)
	if !ok {
		http.Error(w, "token not found in request", http.StatusUnauthorized)
	}

	fmt.Println(token)

}
