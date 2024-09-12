package api

import (
	"encoding/xml"
	"net/http"
)

func (cfg *ApiConfig) TestResponder(w http.ResponseWriter, r *http.Request) {
	Items := cfg.DataFetcher("https://blog.boot.dev/index.xml")
	response, _ := xml.Marshal(Items)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
