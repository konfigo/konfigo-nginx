FROM golang:1.24 AS builder

WORKDIR /app

COPY src/ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o poll main.go

FROM nginx:alpine AS runner

COPY --from=builder /app/poll /usr/local/bin/poll

RUN chmod +x /usr/local/bin/poll

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
