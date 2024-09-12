package api

import (
	"github.com/google/uuid"
	"github.com/lovesupergames/RSSAgregator/internal/common"
	"github.com/lovesupergames/RSSAgregator/internal/database"
	"net/http"
)

func (cfg *ApiConfig) HandleDeleteFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	FeedId, err := uuid.Parse(r.PathValue("feedFollowID"))
	if err != nil {
		common.RespondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	FeedAfterDelete, err := cfg.DB.DeleteFeedFollow(r.Context(), FeedId)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.RespondWithJSON(w, http.StatusOK, FeedAfterDelete)
}
