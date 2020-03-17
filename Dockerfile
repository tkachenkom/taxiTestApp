FROM golang:1.13.7-stretch

WORKDIR $GOPATH/src/github.com/tkachenkom/taxiTestApp/

COPY . .

ENV GO111MODULE=on

RUN go build -o taxi -v ./cmd/main.go