FROM golang:latest as builder

WORKDIR /app
COPY ./ ./
COPY go.mod .
COPY go.sum .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

WORKDIR /app
COPY --from=builder ["/cmd/api_main", "/app"]
FROM alpine:latest
CMD ["./cmd/main"]