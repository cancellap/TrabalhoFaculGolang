FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .   
COPY .env .env
RUN go build -o api cmd/api/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/api .
COPY --from=builder /app/.env . 

EXPOSE 8080

CMD ["./api"]
