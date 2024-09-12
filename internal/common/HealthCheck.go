package common

import "net/http"

type healthz struct {
	Status string `json:"status"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, http.StatusOK, healthz{Status: "OK"})
}
