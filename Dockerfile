FROM golang:1.16-alpine

WORKDIR /go/src/github.com/nht1206/pricetracker

RUN apk add --no-cache git

COPY . /go/src/github.com/nht1206/pricetracker

RUN go get -d -v ./...
RUN go install -v ./...