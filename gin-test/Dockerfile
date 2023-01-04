FROM golang:1.17.7-alpine

WORKDIR /go/src/work
COPY . .

# Air導入
RUN apk update &&  apk add --no-cache git
RUN go get github.com/cosmtrek/air@v1.29.0 && go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]
