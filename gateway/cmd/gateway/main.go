/*
	Gateway Service waits message from client
	When message comes from client new thread is opened and that thread waits for response from service(RPC)
	Main thread still waits for message from client
	When RPC returns at other thread, gateway returns response
*/

package main

import (
	"log"
	"net/http"

	"gateway/pkg/api"
	"gateway/pkg/api/http/httpServer"
)

func main() {
	http.HandleFunc("/", httpServer.GetChangedText)

	api.ConnectService()

	// Serve
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatalf("Gateway Service couldn't listen port 3333: %s", err.Error())
	}

	api.DisconnectService()
}
