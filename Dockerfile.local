FROM golang:1.23.4

RUN apt-get update && apt-get install -y \
    curl \
    && rm -rf /var/lib/apt/lists/*

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY . .


RUN go mod tidy
RUN go build -o main .

RUN chmod +x /app/main

EXPOSE 8000

CMD ["sh", "-c", "air"]
