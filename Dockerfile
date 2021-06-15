FROM golang:alpine AS build-env
ADD . /app
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct
ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0
RUN cd /app && go build -o server


FROM alpine
WORKDIR /app
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && apk del tzdata
ENV APP_RUN_ENV pro
COPY --from=build-env /app/server /app/server
COPY --from=build-env /app/.env.pro /app/.env.pro
ENTRYPOINT ./server
