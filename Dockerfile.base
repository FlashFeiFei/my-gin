FROM golang:1.13

WORKDIR /go/src/FlashFeiFei/my-gin
#设置golang mod的阿里云代理
ENV GOPROXY https://mirrors.aliyun.com/goproxy/

COPY . .

RUN go mod download

RUN go install -v ./...

CMD ["my-gin"]