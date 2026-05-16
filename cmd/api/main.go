package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/MarsuvesVex/go-microservice-template/internal/platform/config"
	httptransport "github.com/MarsuvesVex/go-microservice-template/internal/transport/http"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	log.Printf("api listening on %s", cfg.HTTPAddress)

	log.Fatal(http.ListenAndServe(cfg.HTTPAddress, httptransport.NewRouter()))
}
