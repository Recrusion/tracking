FROM golang:1.24.4

WORKDIR /tracking

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o tracking ./cmd/tracking

EXPOSE 8080

CMD ["./tracking"]