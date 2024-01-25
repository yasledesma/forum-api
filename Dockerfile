FROM golang:1.21-alpine3.19

WORKDIR /usr/src/build

RUN go install github.com/cosmtrek/air@latest && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

COPY go.mod go.sum ./
RUN go mod download && go mod verify

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
