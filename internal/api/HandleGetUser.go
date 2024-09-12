package api

import (
	"github.com/lovesupergames/RSSAgregator/internal/common"
	"github.com/lovesupergames/RSSAgregator/internal/database"
	"net/http"
)

func (cfg *ApiConfig) HandleGetUsers(w http.ResponseWriter, r *http.Request, dbApi database.User) {
	common.RespondWithJSON(w, http.StatusOK, databaseUserToUser(dbApi))
}
