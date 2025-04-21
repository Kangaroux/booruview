FROM golang:1.24-alpine AS build

WORKDIR /app

# change watching
RUN go install github.com/mitranim/gow@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# RUN go build -v -o /dist/server .


# FROM scratch

# COPY --from=build /dist/server /

# CMD ["/server"]
