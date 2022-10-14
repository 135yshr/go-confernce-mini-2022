FROM golang:1.19 as builder

WORKDIR /app
ADD . .
RUN go build -o smacare cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=builder /app/smacare .
CMD ["./smacare"]

