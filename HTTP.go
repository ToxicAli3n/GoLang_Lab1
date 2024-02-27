package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, _ *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)
	response := TimeResponse{Time: currentTime}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/time", timeHandler)
	err := http.ListenAndServe(":8795", nil)
	if err != nil {
		return
	}
}
