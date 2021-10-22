FROM golang:latest as builder
WORKDIR /app
COPY go.* .
COPY .env .
RUN ls -la .

RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go
FROM alpine:latest
WORKDIR /app
COPY --from=builder ["/app/main", "/app"]

CMD ["sh", "-c", "/app/main"]