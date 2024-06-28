package main

import "net/http"

func (cfg *apiConfig) handlerReadiness(w http.ResponseWriter, r *http.Request) {
    type response struct {
        Status string `json:"status"`
    }

    respondWithJSON(w, http.StatusOK, response{
        Status: "ok",
    })
}

func (cfg *apiConfig) handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
