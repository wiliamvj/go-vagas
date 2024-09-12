# Fase de construção
FROM golang:1.23 AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o bot cmd/main.go

FROM golang:1.22 AS runner

WORKDIR /app

COPY --from=builder /app/bot /bot

ENTRYPOINT ["/bot"]
