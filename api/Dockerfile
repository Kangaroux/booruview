FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /dist/server .


FROM scratch

COPY --from=build /dist/server /

CMD ["/server"]
