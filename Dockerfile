FROM golang:alpine
LABEL maintainer="Pavel Sumin"

COPY . /code
WORKDIR /code

RUN apk add git
RUN go get github.com/segmentio/kafka-go
RUN go build -o main

CMD ["./main"]
