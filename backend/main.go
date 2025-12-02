package main

import (
	"log"
	"net/http"

	"masoho/backend/handlers"
)

func main() {
	// ルーティング設定
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/users", handlers.UsersHandler)

	addr := ":8080"
	log.Printf("listening on %s ...", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}


