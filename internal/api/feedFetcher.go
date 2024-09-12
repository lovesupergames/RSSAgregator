package api

import (
	"fmt"
	"github.com/lovesupergames/RSSAgregator/internal/common"
	"net/http"
)

func (cfg *ApiConfig) FeedFetcher(w http.ResponseWriter, r *http.Request) {
	feedStruct, err := cfg.DB.GetNextFeedsToFetch(r.Context(), 1)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	for _, feed := range feedStruct {
		xml := cfg.DataFetcher(feed.Url)
		fmt.Println(xml.Channel.Title)
	}
}
