FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o proxy ./cmd/proxy/proxy.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/proxy .

EXPOSE ${PORT}

CMD ["./proxy"]
