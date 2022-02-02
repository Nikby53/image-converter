FROM public.ecr.aws/docker/library/golang:latest as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

FROM public.ecr.aws/docker/library/alpine:latest
WORKDIR /app
COPY --from=builder ["/app/main", "/app"]
COPY ["/api/openapi-spec/swagger.yaml", "/app/api/openapi-spec/"]
COPY ["./.env", "/app"]
CMD ["sh", "-c", "/app/main"]