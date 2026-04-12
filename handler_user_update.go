package main

import (
	"encoding/json"
	"net/http"

	"github.com/rc5091119-pixel/Chirpy/internal/auth"
	"github.com/rc5091119-pixel/Chirpy/internal/database"
)

func (cfg *apiConfig) handlerUserUpdate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	type response struct {
		User
	}
	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "could't get token", err)
		return
	}
	uuid, err := auth.ValidateJWT(tokenString, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid Token", err)
		return
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "could't decode request", err)
		return
	}
	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could't hashed password", err)
		return
	}
	user, err := cfg.db.UpdateUser(r.Context(), database.UpdateUserParams{
		ID:             uuid,
		Email:          params.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could't update user", err)
		return
	}

	respondWithJSON(w, 200, response{
		User: User{
			ID:          user.ID,
			Email:       user.Email,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
			IsChirpyRed: user.IsChirpyRed,
		},
	})

}
