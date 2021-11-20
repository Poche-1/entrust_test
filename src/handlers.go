package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(write_ http.ResponseWriter, req *http.Request) {
	write_.WriteHeader(http.StatusOK)
}

func CarPostRequest(write_ http.ResponseWriter, req *http.Request) {
	var car Car
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&car)
	write_.Header().Set("Content-Type", "application/json")
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", err)
		return
	}
	err = insertDB(&car)
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", msgDatabase)
		return
	}
	responseCarBody(&car, write_)
}

func CarGetRequest(write_ http.ResponseWriter, req *http.Request) {
	var car Car
	err := getDB(&car, req.Header.Get("id"))
	write_.Header().Set("Content-Type", "application/json")
	if err != nil {
		write_.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", msgNotFound)
		return
	}
	write_.WriteHeader(http.StatusOK)
	responseCarBody(&car, write_)
}

func CarDeleteRequest(write_ http.ResponseWriter, req *http.Request) {
	var car Car
	err := getDB(&car, req.Header.Get("id"))
	write_.Header().Set("Content-Type", "application/json")
	if err != nil {
		write_.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", msgNotFound)
		return
	}
	err = deleteDB(&car)
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", msgDatabase)
		return
	}
	write_.WriteHeader(http.StatusNonAuthoritativeInfo)
	responseCarBody(&car, write_)
}

func responseCarBody(car *Car, write_ http.ResponseWriter) {
	response, err := car.ToJson()
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", msgMalFormat)
		return
	}
	write_.Write(response)
}
