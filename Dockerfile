FROM golang:1.18.4
WORKDIR "/app"

COPY . .

ENV GO111MODULE=on\
    GOPROXY=https://goproxy.cn,direct

RUN go mod tidy
RUN go build -o mys_auto_checkin /app/main.go
ENTRYPOINT ["/app/mys_auto_checkin"]