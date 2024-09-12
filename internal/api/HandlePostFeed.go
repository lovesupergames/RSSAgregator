package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/lovesupergames/RSSAgregator/internal/common"
	"github.com/lovesupergames/RSSAgregator/internal/database"
	"net/http"
	"time"
)

func (cfg *ApiConfig) HandlePostFeed(w http.ResponseWriter, r *http.Request, dbApi database.User) {
	type payload struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	pld := payload{}
	req := json.NewDecoder(r.Body)
	err := req.Decode(&pld)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      pld.Name,
		Url:       pld.Url,
		Userid:    dbApi.ID,
		FeedID:    uuid.New(),
	})
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	common.RespondWithJSON(w, http.StatusOK, databaseFeedToUser(feed))
}
