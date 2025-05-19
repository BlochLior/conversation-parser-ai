FROM golang:1.24.3-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o frontend-server ./frontend

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/frontend-server .
COPY frontend/static/ static/

EXPOSE 8080
CMD ["./frontend-server"]
