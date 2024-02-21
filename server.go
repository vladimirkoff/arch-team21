package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type timeJSON struct {
	CURRENT_TIME string `json:"time"`
}

func getCurrentTime(w http.ResponseWriter, r *http.Request) {
	currentTime := timeJSON{CURRENT_TIME: time.Now().Format(time.RFC3339)}
	json_data, err := json.Marshal(currentTime)
}

func main() {}
