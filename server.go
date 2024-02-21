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

func getTime(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("time request received\n")

	currentTime := timeJSON{CURRENT_TIME: time.Now().Format(time.RFC3339)}
	json_data, err := json.Marshal(currentTime)

	if err != nil {
		log.Fatal(err)
	}

	io.WriteString(w, string(json_data))
}

func main() {
	http.HandleFunc("/time", getTime)
	http.ListenAndServe(":8795", nil)
}
