# Dockerfile para Go CLI
FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o cli main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/cli ./cli
ENTRYPOINT ["./cli"]
