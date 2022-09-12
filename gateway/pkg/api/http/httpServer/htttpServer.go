package httpServer

import (
	"fmt"
	"gateway/pkg/api"
	"io"
	"log"
	"net/http"
)

func GetChangedText(w http.ResponseWriter, r *http.Request) {
	// Read request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Read HTTP request body failed: %s", err.Error())
	}
	fmt.Println("Body read with value: %s", string(body))

	// Pass request wait for response
	serviceResponse := api.SendMessageToService(string(body))
	fmt.Println("Service response done")

	// Return response
	w.Write([]byte(serviceResponse))
	fmt.Println("HTTP response done")
	return
}
