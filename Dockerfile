FROM golang:1.16

WORKDIR /go/src/one-direct
COPY . .
RUN go get
RUN go build -o main *.go

ENV PORT=80

EXPOSE 80/tcp
EXPOSE 80/udp

ENTRYPOINT ["/go/src/one-direct/main"]
