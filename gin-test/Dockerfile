FROM golang:1.17.7-alpine

RUN mkdir /go/src/work
WORKDIR /go/src/work

# Air導入
RUN apk update &&  apk add git
RUN go get github.com/cosmtrek/air@v1.29.0
CMD ["air", "-c", ".air.toml"]

ADD . /go/src/work