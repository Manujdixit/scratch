package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code>499{
		log.Printf("Error: %s",message)}
		type errResponse struct{
			Error string `json:"error"`
		}
		respondWithJson(w, code, errResponse{message})
	}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	dta,err:=json.Marshal(payload)
	if err!=nil{
		log.Printf("failed  to marshal: %v",payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(dta)
}