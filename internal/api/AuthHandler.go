package api

import (
	"github.com/lovesupergames/RSSAgregator/internal/common"
	"github.com/lovesupergames/RSSAgregator/internal/database"
	"net/http"
	"strings"
)

type AuthedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) MiddlewareAuthUser(handler AuthedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getApi := r.Header.Get("Authorization")
		getApi = strings.TrimPrefix(getApi, "ApiKey ")
		if getApi == "" {
			common.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		dbApi, err := cfg.DB.GetUserByApi(r.Context(), getApi)
		if err != nil {
			common.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// Call the next handler on successful authentication, passing `dbApi` as the user
		handler(w, r, dbApi)
	}
}
