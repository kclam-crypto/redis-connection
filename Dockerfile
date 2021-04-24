FROM golang:1.15 AS BUILDER
WORKDIR /app
COPY . .
RUN go get github.com/gomodule/redigo/redis
RUN go build main.go

FROM alpine:3
WORKDIR "/app"
RUN mkdir "/lib64" && ln -s "/lib/libc.musl-x86_64.so.1" "/lib64/ld-linux-x86-64.so.2"
COPY --from=builder "/app/main" .
CMD ["sleep", "infinity"]
