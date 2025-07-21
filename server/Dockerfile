FROM golang:alpine as builder

WORKDIR /stock-exchange/server
COPY . .

RUN go env -w GO111MODULE=on \
&& go env -w GOPROXY=https://goproxy.cn,direct \
&& go env -w CGO_ENABLED=0 \
&& go env \
&& go mod tidy \
&& go build -o server .

FROM alpine:latest

WORKDIR /stock-exchange/server

COPY --from=0 /stock-exchange/server/server ./
COPY --from=0 /stock-exchange/server/resource ./resource/
COPY --from=0 /stock-exchange/server/config.docker.yaml ./

RUN apk add --no-cache tzdata \
&& ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Asia/Shanghai" > /etc/timezone

EXPOSE 8888
ENTRYPOINT ./server -c config/config.yaml
