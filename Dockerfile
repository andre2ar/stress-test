FROM golang:latest

WORKDIR /app

COPY . .

RUN GOOS=linux go build .

ENTRYPOINT ["./stress-test", "stress-test"]