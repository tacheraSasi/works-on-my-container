# works-on-my-container

> A Docker playground where everything is in containers, and your laptop silently judges you.

## What’s inside?

* Go API (`/api`)
* PostgreSQL (`/db`)
* Redis (`/redis`)
* Nginx reverse proxy (`/nginx`)

All wired together in **docker-compose**. Because apparently, everything works better in a container.

## Quick Start

```bash
docker compose up --build
```

Then open [http://localhost](http://localhost) and enjoy the chaos.

## Notes

* API logs go to `api` container
* DB is persistent thanks to volumes
* Nginx is just there to look fancy
