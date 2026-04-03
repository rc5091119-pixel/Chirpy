package main

import (
	"encoding/json"
	"net/http"

	"github.com/rc5091119-pixel/Chirpy/internal/auth"
)
func (cfg *apiConfig) handlerLogin(w http.ResponseWriter,r *http.Request){
	type parameters struct{
		Password string `json:"password"`
		Email string `json:"email"`
	}
	params := parameters{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	
	user,err := cfg.db.GetUser(r.Context(),params.Email)

	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", nil)
		return
	}

	check,err := auth.CheckPasswordHash(params.Password,user.HashedPassword)
	if err != nil {
		respondWithError(w,http.StatusInternalServerError,"Incorrect email or password",err)
		return
	}
	if !check {
		respondWithError(w,http.StatusUnauthorized,"Incorrect email or password",nil)
		return
	}
	login := User{
		ID: user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email: user.Email,
	}
	
	respondWithJSON(w,http.StatusOK,login)
}