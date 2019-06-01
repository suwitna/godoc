FROM golang:latest

RUN mkdir -p /go/src/godoc
ADD . /go/src/godoc

WORKDIR /go/src/godoc

RUN go build -o /go/bin/godoc
CMD ["/go/bin/godoc"]
