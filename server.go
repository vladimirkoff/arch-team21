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

func getTime(writer http.ResponseWriter, r *http.Request) {
	fmt.Printf("time request received\n")

	currentTime := timeJSON{CURRENT_TIME: time.Now().Format(time.RFC3339)}
	data, err := json.Marshal(currentTime)

	if err != nil {
		log.Fatal(err)
	}

	io.WriteString(writer, string(data))
}

func main() {
	http.HandleFunc("/time", getTime)
	http.ListenAndServe(":8795", nil)
}
