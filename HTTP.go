package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)
	res := TimeResponse{Time: currentTime}
	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/time", timeHandler)
	err := http.ListenAndServe(":8795", nil)
	if err != nil {
		panic(err)
	}
}
