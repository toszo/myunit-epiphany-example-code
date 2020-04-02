FROM golang

COPY main.go .

RUN go build



FROM alpine

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY --from=0 /go/go .

EXPOSE 8080

ENTRYPOINT /go