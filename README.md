# booruview

[booruview.com](https://booruview.com) is a web app for browsing gelbooru posts. It's very fast, has a simple interface, and anonymous with no signups or search limits.

### Running locally

Start the dev server

```bash
./dev.sh

# or
docker compose up --build
```

Then visit http://localhost:8080

### Infrastructure

The site uses three components:

- API built with Go. Proxies requests to the gelbooru API, transforms the response, and caches it
- Frontend built with Vue.js and Typescript
- Valkey (Redis) as the backend cache

It also includes [Loki](https://grafana.com/oss/loki/) and [Grafana](https://grafana.com/oss/grafana/) for monitoring.

