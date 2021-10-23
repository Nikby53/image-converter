FROM golang:latest as builder
WORKDIR /api

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go
FROM alpine:latest
COPY --from=builder ["/api/main", "/"]

ENTRYPOINT ["/main"]