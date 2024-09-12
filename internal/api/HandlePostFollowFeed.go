package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/lovesupergames/RSSAgregator/internal/common"
	"github.com/lovesupergames/RSSAgregator/internal/database"
	"net/http"
)

func (cfg *ApiConfig) HandlePostFollowFeed(w http.ResponseWriter, r *http.Request, feed database.User) {
	type payload struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	pld := payload{}
	req := json.NewDecoder(r.Body)
	if err := req.Decode(&pld); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}
	followFeed, err := cfg.DB.CrateFeedFollow(r.Context(), pld.FeedId)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	common.RespondWithJSON(w, http.StatusOK, followFeed)
}
