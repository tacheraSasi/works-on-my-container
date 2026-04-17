# works-on-my-container

> My Docker playground where everything is in containers, and my laptop humbles me.

## What's inside?

* **Gateway** — HTTP API that pretends to know what it's doing, then asks other services via gRPC
* **Users Service** — gRPC service that returns users.
* **PostgreSQL** — a real database for when JSON files stop being funny
* **Redis** — sitting there, waiting to be useful
* **Nginx** — reverse proxy. Still just here to look fancy
* **Protobuf** — because REST is for people who like reading JSON at 3am

## Project Structure

```
├── Makefile              # the one file that actually runs things
├── docker-compose.yml    # the orchestrator of chaos
├── go.mod                # single module, no multi-module drama
├── proto/                # protobuf definitions + generated code
│   └── user.proto
├── services/
│   ├── gateway/          # HTTP gateway (port 8080)
│   └── users/            # gRPC users service (port 8081)
├── shared/               # shared utilities across services
├── db/
│   └── init.sql          # DB init script
└── nginx/
    └── nginx.conf        # nginx config
```

## Quick Start

```bash
make up
```

This regenerates protobuf code and runs `docker compose up -d --build`. Because we're civilized.

To tear it all down:

```bash
make down
```

Then open [http://localhost](http://localhost) and enjoy the chaos.

## Other Commands

| Command | What it does |
|---------|-------------|
| `make proto` | Regenerate protobuf Go code |
| `make up` | Proto + build + start everything |
| `make down` | Stop and remove containers |

## Prerequisites

* Docker & Docker Compose
* Go 1.24+
* `protoc` with `protoc-gen-go` and `protoc-gen-go-grpc` plugins (for proto regeneration)

## Notes

* Gateway talks to Users Service over gRPC (insecure, because they're on the same Docker network and trust is earned through proximity)
* DB is persistent thanks to volumes
* Proto files live at the repo root so both services can import the generated code without multi-module suffering
* The module is called `works-on-my-machine` because honesty matters
