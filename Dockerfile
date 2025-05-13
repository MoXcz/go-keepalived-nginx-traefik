FROM golang:1.24-alpine

WORKDIR /app

COPY . .

RUN go build -o server .

EXPOSE 4000

CMD ["./server", "-addr=:4000"]

