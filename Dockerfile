FROM golang:1.15 AS BUILDER
WORKDIR /app
COPY . .
RUN go get github.com/gomodule/redigo/redis
RUN go build main.go

FROM ubuntu:20.04
WORKDIR "/app"
COPY --from=builder "/app/main" .
CMD ["sleep", "infinity"]
