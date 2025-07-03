package main

import (
	"log"
	"sso-go-gin/config"
)

func main() {
	cfg := config.Load()

	engine, err := InitializeApp(cfg)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	if err := engine.RunTLS(":8080", "./certs/localhost+2.pem", "./certs/localhost+2-key.pem"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
