FROM golang:1.8
RUN apt-get update -qq && \
    apt-get install -y mysql-client vim

WORKDIR /go/src/go_line_bot

ADD . .
RUN go build main.go
