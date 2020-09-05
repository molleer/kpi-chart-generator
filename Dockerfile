FROM golang:1.5 AS build
WORKDIR /go/src/action
COPY . .
RUN go get -d -v
RUN go install -v
RUN go build .