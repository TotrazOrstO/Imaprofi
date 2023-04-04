FROM golang:1.17

WORKDIR /app

COPY go.mod go.sum .env ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

CMD ["/app/main"]