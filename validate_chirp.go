package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (cfg *apiConfig)handlerValidate_chrip(w http.ResponseWriter, r *http.Request) {
	type chirpRequest struct {
		Body string `json:"body"`
	}
	// type validResponse struct {
	// 	Valid bool `json:"valid"`
	// }
	type cleanedBody struct{
		Cleanedbody string `json:"cleaned_body"`
	}
	decoder := json.NewDecoder(r.Body)
	chriprequest := chirpRequest{}
	err := decoder.Decode(&chriprequest)
	if err != nil {
		respondeWithError(w,http.StatusInternalServerError,"could't decode parameter",err)
		return
	}
	s := chriprequest.Body

	temp := strings.Split(s," ")
	
	
	for i,v := range temp{
		if strings.ToLower(v) == "kerfuffle" || strings.ToLower(v) == "sharbert" || strings.ToLower(v) == "fornax"{
		temp[i] = "****"
	    }
    }
	s = strings.Join(temp," ")
	
	if len(s) > 140 {
		respondeWithError(w,http.StatusBadRequest,"Chirp is too long",nil)
		return
	}

	respondeWithJosn(w,http.StatusOK,cleanedBody{
		Cleanedbody: s,
	})
}
