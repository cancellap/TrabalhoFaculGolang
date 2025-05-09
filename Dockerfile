FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .   
RUN go build -o api main.go

# Etapa final
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/api .
COPY --from=builder /app/.env . 

EXPOSE 8080

CMD ["./api"]
