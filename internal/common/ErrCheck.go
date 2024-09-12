package common

import "net/http"

type ErrChecker struct {
	ErrChk string `json:"error"`
}

func ErrCheck(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
