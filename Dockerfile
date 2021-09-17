FROM golang:alpine AS build-env
ADD . /app
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct
ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0
RUN cd /app \
    && go mod download \
    && time go build -o goal


FROM alpine
WORKDIR /app
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && apk del tzdata
COPY --from=build-env /app/goal /app/goal
COPY --from=build-env /app/.env /app/.env
ENTRYPOINT ./goal
