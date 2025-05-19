# use compatible toolchain to avoid go.mod go version mismatch
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

# currently, no dependencies exist. when they do, update this code to `COPY go.mod go.sum ./`
COPY . .
# ignore if no external dependencies:
RUN go build -o server ./go-backend

# --- Runtime stage ---
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server ./

EXPOSE 8000
CMD ["./server"]