FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./
RUN go build -o image-converter ./cmd/main.go

CMD ["./image-converter"]