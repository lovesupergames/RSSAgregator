package api

import (
	"github.com/lovesupergames/RSSAgregator/internal/common"
	"net/http"
)

func (cfg *ApiConfig) HandleGetAllFeed(w http.ResponseWriter, r *http.Request) {
	allFeeds, err := cfg.DB.GetAllFeed(r.Context())
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.RespondWithJSON(w, http.StatusOK, allFeeds)
}
