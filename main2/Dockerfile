FROM golang:1.14-alpine AS build

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]