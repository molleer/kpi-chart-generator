FROM golang:latest
WORKDIR /go/src/action
COPY . .
RUN go get -d -v
RUN go install -v
RUN go build main.go
CMD ["./main"]
