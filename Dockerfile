FROM golang:1.16.7-alpine
RUN apk add build-base

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN mkdir ./data
RUN mkdir ./server
COPY data/*.go ./data
COPY server/*.go ./server
WORKDIR /app/server
RUN go build -ldflags="-w -s" -o server

EXPOSE 3000
CMD ["./server", "start"]