FROM golang:1.10 as builder

WORKDIR /go/src/github.com/vnkrtv/go-habr-loader
COPY . .
RUN go get -t .../ \
 && go build -ldflags "-linkmode external -extldflags -static" -a cmd/main.go

FROM alpine:3.6 as alpine
RUN apk add -U --no-cache ca-certificates

FROM scratch
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/vnkrtv/go-habr-loader/main /main
COPY config /config
CMD ["/main"]
