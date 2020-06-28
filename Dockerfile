FROM golang:1.13-alpine

RUN apk add git; \
    apk add curl; \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mkdir -p /go/src/github.com/okapusta/employee-records-go

WORKDIR /go/src/github.com/okapusta/employee-records-go

COPY . .

RUN dep ensure

RUN go build -o bin/employees ./cmd/employees

RUN ln -s `pwd`/bin/employees /go/bin; \
    chmod a+x /go/bin/employees
