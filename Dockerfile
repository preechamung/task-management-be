FROM golang:1.16-alpine AS builder
WORKDIR /app
COPY . .
RUN GIN_MODE=release
RUN go mod download
RUN go build -o ./build ./cmd/main.go

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/build .
COPY --from=builder /app/.env .
EXPOSE ${CONTAINER_PORT}
ENTRYPOINT ["./build"]