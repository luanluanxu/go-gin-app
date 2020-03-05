FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/luanluanxu/go-gin-app/my_app
COPY . $GOPATH/src/github.com/luanluanxu/go-gin-app
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-app"]
