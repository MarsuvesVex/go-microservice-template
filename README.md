# Go Microservice Template

<div align="center">

# ⚡ Go Microservice Template

Production-ready Go microservice template for event-driven platforms, SaaS backends, orchestration systems, and scalable internal tooling.

<p>
  <img src="https://img.shields.io/badge/Go-1.22-00ADD8?style=for-the-badge&logo=go" alt="Go" />
  <img src="https://img.shields.io/badge/Postgres-16-336791?style=for-the-badge&logo=postgresql" alt="Postgres" />
  <img src="https://img.shields.io/badge/NATS-JetStream-27AAE1?style=for-the-badge" alt="NATS" />
  <img src="https://img.shields.io/badge/Redis-7-DC382D?style=for-the-badge&logo=redis" alt="Redis" />
  <img src="https://img.shields.io/badge/Docker-Compose-2496ED?style=for-the-badge&logo=docker" alt="Docker" />
  <img src="https://img.shields.io/badge/Traefik-Optional-24A1C1?style=for-the-badge&logo=traefikproxy" alt="Traefik" />
</p>

<p>
  Built for rapidly bootstrapping scalable Go backends with clean architecture,
  event-driven workflows, container-first development, and composable infrastructure layers.
</p>

</div>

---

# ✨ Features

## Core Architecture

* API service
* Worker service
* Scheduler service
* Event-driven architecture
* NATS JetStream event bus
* Structured internal package layout
* Analytics/event ingestion foundation
* Generic service abstractions
* Container-first development workflow

## Infrastructure

* Docker Compose layering
* Internal Nginx reverse proxy
* Optional Traefik overlay
* Postgres
* Redis
* NATS JetStream
* Environment-based configuration
* Isolated internal Docker network

## Developer Experience

* Fast local development
* Clean service separation
* Ready for monorepo workflows
* Generic reusable architecture
* Easy template customization
* Extensible service structure
* Scalable Compose overlays

---

# 🧱 Architecture

```txt
                        ┌───────────────────────┐
                        │       Traefik        │
                        │     (optional)       │
                        └──────────┬────────────┘
                                   │
                        ┌──────────▼───────────┐
                        │        Nginx         │
                        │ Internal Reverse API │
                        └──────────┬───────────┘
                                   │
                 ┌─────────────────┼─────────────────┐
                 │                 │                 │
        ┌────────▼───────┐ ┌──────▼──────┐ ┌────────▼────────┐
        │      API       │ │    Worker   │ │    Scheduler    │
        │ HTTP/REST      │ │ Event Tasks │ │ Cron/Timed Jobs │
        └────────┬───────┘ └──────┬──────┘ └────────┬────────┘
                 │                 │                 │
                 └─────────────────┼─────────────────┘
                                   │
                        ┌──────────▼───────────┐
                        │     NATS JetStream   │
                        │      Event Bus       │
                        └──────────┬───────────┘
                                   │
                    ┌──────────────┼──────────────┐
                    │                             │
           ┌────────▼────────┐         ┌──────────▼─────────┐
           │    Postgres     │         │       Redis        │
           │ Persistence     │         │ Cache / Queues     │
           └─────────────────┘         └────────────────────┘
```

---

# 📁 Project Structure

```txt
.
├── cmd/
│   ├── api/
│   ├── worker/
│   └── scheduler/
│
├── internal/
│   ├── agents/
│   ├── app/
│   ├── connectors/
│   ├── domain/
│   ├── events/
│   ├── platform/
│   ├── services/
│   ├── store/
│   └── transport/
│
├── docker/
│   ├── api/
│   ├── worker/
│   ├── scheduler/
│   └── nginx/
│
├── migrations/
├── docs/
├── scripts/
└── pkg/
```

---

# 🚀 Quick Start

## 1. Clone

```bash
git clone https://github.com/MarsuvesVex/go-microservice-template.git
cd go-microservice-template
```

## 2. Configure

```bash
cp .env.example .env
```

## 3. Install Dependencies

```bash
go mod tidy
```

## 4. Start Infrastructure + Services

```bash
make stack
```

## 5. Run API Locally

```bash
make dev
```

---

# 🔌 Compose Layers

One of the main goals of this template is composable infrastructure.

## Base Infrastructure

```bash
make stack
```

Starts:

* Postgres
* Redis
* NATS
* API
* Worker
* Scheduler

---

## Internal Nginx Layer

```bash
make nginx
```

Adds:

* Internal reverse proxy
* Unified API entrypoint
* Internal routing layer

Health check:

```bash
curl http://localhost:8080/healthz
```

---

## Traefik Layer

Create external Traefik network:

```bash
docker network create traf-external
```

Then:

```bash
make traefik
```

Adds:

* Traefik labels
* Automatic routing
* TLS-ready service exposure
* Internal/external network separation

Default domain:

```txt
go-template.dev.local
```

---

# ⚙️ Environment Variables

```env
SERVICE_NAME=go-microservice-template
HTTP_ADDRESS=:3002
LOG_LEVEL=debug

POSTGRES_DB=app
POSTGRES_USER=app
POSTGRES_PASSWORD=app
POSTGRES_PORT=5434
DATABASE_URL=postgres://app:app@localhost:5434/app?sslmode=disable

REDIS_PORT=6380
REDIS_ADDR=localhost:6380

NATS_PORT=4222
NATS_MONITOR_PORT=8222
NATS_URL=nats://localhost:4222

NGINX_PORT=8080

DOMAIN=go-template.dev.local
TRAEFIK_ROUTER=go-template
TRAEFIK_ENTRYPOINTS=web,websecure
TRAEFIK_CERTRESOLVER=myresolver
TRAEFIK_NETWORK=traf-external
```

---

# 🧠 Designed For

This template is intentionally generic and reusable.

## Ideal Use Cases

* SaaS platforms
* Internal tooling
* Automation systems
* Event-driven backends
* Agent orchestration systems
* AI infrastructure
* Analytics/event ingestion systems
* Queue-based workflows
* Creator tools
* API-first platforms

---

# 📨 Event Bus Foundation

The template includes a generic event abstraction layer.

Example event:

```go
package events

type Event struct {
    ID       string
    Topic    string
    Source   string
    Data     []byte
}
```

Example topics:

```txt
user.created
email.send.requested
analytics.event.ingested
billing.subscription.updated
job.completed
```

---

# 🐳 Docker Philosophy

The infrastructure is intentionally layered.

This allows projects to:

* run locally without Traefik
* use internal-only networking
* scale into shared reverse proxy environments
* keep infrastructure reusable across projects
* compose environments incrementally

---

# 🧩 Template Customization

## Rename Module

Update:

```txt
go.mod
```

Replace imports:

```bash
find . -name "*.go" -type f -exec sed -i 's|github.com/MarsuvesVex/go-microservice-template|github.com/YOUR_USER/YOUR_REPO|g' {} +
```

Update:

* service names
* domains
* Docker Compose project name
* Traefik router labels
* environment variables

---

# 📚 Included Foundations

## HTTP

* Chi router
* Structured handlers
* Service separation
* Transport abstraction

## Persistence

* Postgres
* pgx
* migration structure
* repository/store layer

## Messaging

* NATS JetStream
* event abstraction
* worker architecture

## Infra

* Compose layering
* Nginx reverse proxy
* Traefik overlays
* reusable Dockerfiles

---

# 🛠 Planned Extensions

Potential future additions:

* OpenTelemetry
* Prometheus metrics
* Grafana dashboards
* JWT auth middleware
* Service discovery
* Kubernetes manifests
* GitHub Actions CI
* CQRS helpers
* Generic queue consumers
* gRPC support
* WebSocket gateway
* Distributed tracing

---

# 📄 License

MIT

---

<div align="center">

Built for scalable Go systems ⚡

</div>
