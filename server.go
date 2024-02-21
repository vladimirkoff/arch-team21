package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type timeJSON struct {
	CURRENT_TIME string `json:"time"`
}

func getCurrentTime(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("time request received\n")

	currentTime := timeJSON{CURRENT_TIME: time.Now().Format(time.RFC3339)}
	json_data, err := json.Marshal(currentTime)

	if err != nil {
		log.Fatal(err)
	}

	
}

func main() {}
