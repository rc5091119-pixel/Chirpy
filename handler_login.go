package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rc5091119-pixel/Chirpy/internal/auth"
	"github.com/rc5091119-pixel/Chirpy/internal/database"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	type response struct {
		User
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
	params := parameters{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	user, err := cfg.db.GetUser(r.Context(), params.Email)

	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", nil)
		return
	}

	check, err := auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Incorrect email or password", err)
		return
	}
	if !check {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", nil)
		return
	}
	const days = 60
	refreshTime := time.Hour * 24 * days
	expirationTime := time.Hour
	accessToken, err := auth.MakeJWT(
		user.ID, cfg.jwtSecret,
		expirationTime,
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create access JWT", err)
		return
	}
	refreshTokenstring := auth.MakeRefreshToken()
	_,err = cfg.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token:     refreshTokenstring,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(refreshTime),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create access JWT", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
			IsChirpyRed: user.IsChirpyRed,
		},

		Token:        accessToken,
		RefreshToken: refreshTokenstring,
	})
}
