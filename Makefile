-include .env
export

COMPOSE=docker compose
COMPOSE_APP=$(COMPOSE) -f docker-compose.yml -f docker-compose.services.yml
COMPOSE_NGINX=$(COMPOSE_APP) -f docker-compose.nginx.yml
COMPOSE_TRAEFIK=$(COMPOSE_NGINX) -f docker-compose.traefik.yml

dev:
	go run ./cmd/api

stack:
	$(COMPOSE_APP) up -d --build

nginx:
	$(COMPOSE_NGINX) up -d --build

traefik:
	$(COMPOSE_TRAEFIK) up -d --build

down:
	$(COMPOSE_TRAEFIK) down

fmt:
	go fmt ./...

test:
	go test ./...

tidy:
	go mod tidy
