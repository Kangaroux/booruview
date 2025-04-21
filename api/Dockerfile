FROM golang:1.24-alpine AS build

WORKDIR /app

# change watching
RUN go install github.com/mitranim/gow@latest

# sql migrations
RUN go install \
    -tags="no_clickhouse no_libsql no_mssql no_mysql no_sqlite3 no_vertica no_ydb" \
    github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# RUN go build -v -o /dist/server .


# FROM scratch

# COPY --from=build /dist/server /

# CMD ["/server"]
