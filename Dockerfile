FROM golang:1.24.4

RUN apt-get update && apt-get install -y \
    postgresql-client \
    netcat-openbsd

WORKDIR /tracking

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN go build -o tracking ./cmd/tracking

EXPOSE 8080

CMD ["./tracking"]