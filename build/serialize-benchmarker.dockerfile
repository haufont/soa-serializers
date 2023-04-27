FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o serialize-benchmarker ./cmd/serialize-benchmarker/serialize-benchmarker.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/serialize-benchmarker .

EXPOSE ${PORT}

CMD ["./serialize-benchmarker"]
