# Build stage
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

# Install git and ca-certificates (optional if you use private modules)
RUN apk update && apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/main .

# Render will inject environment variables directly into this container
EXPOSE 8000
CMD ["./main"]
