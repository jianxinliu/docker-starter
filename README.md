# docker app lifecycle

## 使用 go-zero 创建一个简单的 api 程序

```shell
go mode init docker-starter
goctl api new app
go mod tidy
```

## 生成 Dockerfile

```shell
cd app
goctl docker --port 8888
```

## 基于生成的 dockerfile 进行调整

```dockerfile
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
```

## 创建 Makefile

```makefile
BUILD_DEST_DIR ?= build

docker-build:
	docker buildx build -f ./app/Dockerfile -t docker-starter:001 . 

build-app:
	go build -ldflags="-s -w" -o ./{BUILD_DEST_DIR}/docker-starter ./app

.PHONY: docker-build build-app
```

## 构建 docker 镜像

```shell
make docker-build
```

`docekr images` 即可看到生成的镜像

## 运行 docker

```shell
docker run -itd --name docker-starter -p 8888:8888 docker-starter:001

# 查看当前运行的容器
docker container ls -a
```

## 调用 app 接口

```shell
curl http://localhost:8888/for/you
> {"message":"you"}
```

## 查看容器内 app 打印的日志

```shell
# 获取容器 id
docker container ls -a

# 查看容器日志
docker logs <container id>
```

## 查看容器内文件

```shell
docker exec -it <container id> ls -al

```

## 改动容器内容

```shell
# alpine 中的 shell 叫 ash
docker exec it <container id> ash 

# 进入容器内部, 修改配置
vi ./etc/app-api.yaml

# 保存后退出容器，查看变更，会列出容器内文件的变更
docker diff <container id>

```