# golang 作为基础镜像
FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

ADD go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod,sharing=locked \
    go mod download
COPY . .
# 构建 go app, 并将构建结果，写到容器 build 目录下
RUN go build -ldflags="-s -w" -o ./docker-starter ./app

# 构建配置文件 & 复制到容器内
FROM bhgedigital/envsubst as config
ARG ENVFILE=preview.env
WORKDIR /build
COPY ./app/etc/templates/ ./etc
RUN export $(cat ./etc/${ENVFILE} | xargs) && \
    envsubst < ./etc/config.tmpl >> ./etc/app-api.yaml

# 将应用跑在 alpine 上
FROM alpine:3.16

WORKDIR /app
COPY --from=config /build/etc/app-api.yaml ./etc/app-api.yaml
COPY --from=builder /build/docker-starter ./
RUN chmod +x ./docker-starter

EXPOSE 8888
ENTRYPOINT ["/app/docker-starter"]
