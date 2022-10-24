FROM golang:1.19-alpine AS builder

WORKDIR /src

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY .. ./
RUN go build -ldflags "-s -w" -o "./bin/server" "./cmd/server"

FROM alpine:latest
COPY --from=builder /src/bin/server /app/server

WORKDIR /app
EXPOSE 8080

CMD ["/app/server"]