FROM golang
RUN go get github.com/couchbaselabs/n1qlgen && \
    go build github.com/couchbaselabs/n1qlgen
ENTRYPOINT ["/go/bin/n1qlgen"]
