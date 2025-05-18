FROM golang:1.24.2-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o frontend-server ./main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/frontend-server .
COPY static/ static/

EXPOSE 8080
CMD ["./frontend-server"]
