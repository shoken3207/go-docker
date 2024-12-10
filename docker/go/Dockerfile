FROM golang:1.23.3-alpine
WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN go install github.com/air-verse/air@latest
CMD ["air"]