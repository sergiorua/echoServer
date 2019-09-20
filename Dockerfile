FROM golang:1.12 as builder

RUN mkdir -p $GOPATH/src/github.com/sergiorua/echoserver
ADD . $GOPATH/src/github.com/sergiorua/echoserver
WORKDIR $GOPATH/src/github.com/sergiorua/echoserver

RUN go get -d -v ./... && go install -v ./...

RUN go build echoserver.go && ls -l && echo $GOPATH

FROM golang:1.12
RUN mkdir /app
COPY --from=builder /go/src/github.com/sergiorua/echoserver/echoserver /app/
WORKDIR /app
CMD ["./echoserver"]
