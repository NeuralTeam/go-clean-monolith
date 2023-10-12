FROM golang:1.21-bullseye

WORKDIR /opt/backend

COPY . /opt/backend

RUN go mod tidy
CMD go run cmd/main.go start