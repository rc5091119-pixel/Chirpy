package main

import (
	"net/http"
	"time"

	"github.com/rc5091119-pixel/Chirpy/internal/auth"
)
func (cfg *apiConfig)handlerRefresh(w http.ResponseWriter,r *http.Request){
	type response struct{
		Token string `json:"token"`
	}
	tokenstring,err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w,401,"could't get token",err)
		return
	}
	user,err := cfg.db.GetUserFromRefreshToken(r.Context(),tokenstring)
	if err != nil {
		respondWithError(w,401,"could't get token",err)
		return
	}
	accessToken,err := auth.MakeJWT(user.ID,cfg.jwtSecret,time.Hour)
	if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Couldn't create access token", err)
        return
    }
	respondWithJSON(w,200,response{
		Token: accessToken,
	})
}