FROM golang:alpine as builder

ENV ENV GO111MODULE=on \
    CGO_ENABLED=0 

WORKDIR /go/src/oneDirect/
COPY *.go go.mod ./
RUN go mod download
RUN go build -o main *.go

# runtime
FROM alpine
ENV PORT 80

WORKDIR /root/
COPY assets/ ./assets/
COPY *.html *.svg ./
COPY --from=builder /go/src/oneDirect/main ./
# RUN apk add --no-cache ca-certificates

CMD ["/root/main"]
