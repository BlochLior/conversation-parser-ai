# use compatible toolchain to avoid go.mod go version mismatch
FROM golang:1.22-alpine AS builder

WORKDIR /app

# currently, no dependencies exist. when they do, update this code to `COPY go.mod go.sum ./`
COPY go.mod ./
# ignore if no external dependencies:
RUN go mod download || true 

COPY . .
RUN GOTOOLCHAIN=go1.24.2 go build -o server

# --- Runtime stage ---
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/server ./

EXPOSE 8000
CMD ["./server"]