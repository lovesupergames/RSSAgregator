package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/lovesupergames/RSSAgregator/internal/common"
	"github.com/lovesupergames/RSSAgregator/internal/database"
	"net/http"
	"time"
)

func (cfg *ApiConfig) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	type payload struct {
		Name string `json:"name"`
	}
	pld := payload{}
	req := json.NewDecoder(r.Body)
	err := req.Decode(&pld)
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      pld.Name,
	})
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	common.RespondWithJSON(w, http.StatusOK, databaseUserToUser(user))

}
