package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondeWithError(w http.ResponseWriter,code int,msg string,err error){
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	type errorResponse struct{
		Error string `json:"error"`
	}
	respondeWithJosn(w,code,errorResponse{
		Error: msg,
	})
}

func respondeWithJosn(w http.ResponseWriter,code int,payload interface{}){
	w.Header().Set("content-type","application/json")
	dat,err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}