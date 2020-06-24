# Install go-lang for Linux
# A workspace (GOPATH) configured at /go.
FROM golang:1.14-alpine

RUN apk update

COPY ./myweb.go /

WORKDIR /

RUN go build myweb.go

CMD [ "./myweb","&" ]

RUN if [ "$PORT" == "" ]; then export PORT=8080; fi

ENV port=$PORT

EXPOSE ${port}

