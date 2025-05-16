FROM caddy:2.10-builder-alpine AS builder

# transform plugin for formatting fail2ban logs
RUN xcaddy build --with github.com/caddyserver/transform-encoder@15eef974326185626c6d33e5e5f7f452dc6518b5

FROM caddy:2.10-alpine

COPY --from=builder /usr/bin/caddy /usr/bin/caddy
COPY caddy/Caddyfile /etc/caddy/Caddyfile
COPY client/dist/ /srv
