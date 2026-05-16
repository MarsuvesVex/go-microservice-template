# Go Microservice Template

Generic Go microservice template with:

- API service
- Worker service
- Scheduler service
- Postgres
- Redis
- NATS
- Internal Nginx layer
- Optional Traefik layer
- Event bus abstraction
- Docker compose layering

## Setup

```bash
cp .env.example .env
go mod tidy
make stack
make dev
````

## Internal Nginx

```bash
make nginx
```

## Traefik

```bash
docker network create traf-external
make traefik
```

