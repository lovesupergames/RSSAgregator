package api

import (
	"github.com/lovesupergames/RSSAgregator/internal/common"
	"github.com/lovesupergames/RSSAgregator/internal/database"
	"net/http"
)

func (cfg *ApiConfig) HandleGetAllFeedForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	AllFeedForUser, err := cfg.DB.GetAllFeedForUser(r.Context(), user.ID)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.RespondWithJSON(w, http.StatusOK, AllFeedForUser)
}
