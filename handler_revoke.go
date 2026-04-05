package main

import (
	"net/http"

	"github.com/rc5091119-pixel/Chirpy/internal/auth"
)
func(cfg *apiConfig) handlerRevoke(w http.ResponseWriter,r *http.Request){
	s,err := auth.GetBearerToken(r.Header)
	if err != nil {
        respondWithError(w, http.StatusBadRequest, "Couldn't find token", err)
        return
    }
	_,err = cfg.db.RevokeRefreshToken(r.Context(),s)
	if err != nil {
		respondWithError(w,http.StatusInternalServerError,"could't revoke token",err)
		return 
	}
	w.WriteHeader(http.StatusNoContent)
}