FROM golang:latest as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder ["/app/main", "/app"]
COPY ["/api/openapi-spec/swagger.yaml", "/app"]
COPY ["./.env", "/app"]
CMD ["sh", "-c", "/app/main"]