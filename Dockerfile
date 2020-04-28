FROM golang:1.13.5-alpine AS builder

ENV GOPATH $GOPATH:/go
ENV PATH $PATH:$GOPATH/bin

RUN apk update && \
  apk add --no-cache git && \

ADD . /go/src/github.com/todo-app
WORKDIR /go/src/github.com/todo-app
EXPOSE 8000

RUN go build -o main main.go

FROM alpine:3.9
RUN apk add --no-cache tzdata
COPY --from=builder /go/src/github.com/todo-app .
