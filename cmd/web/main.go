package main

import (
	"log"

	"github.com/Ashkan0026/websockets/chat-app/handlers"

	"net/http"

	"github.com/bmizerany/pat"
)

func main() {
	route := Routes()

	log.Println("starting listener channel")

	go handlers.ListenToWsChannel()

	log.Println("Starting server on port 8080")

	_ = http.ListenAndServe(":8080", route)
}

func Routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))
	return mux
}
