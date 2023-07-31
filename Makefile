BUILD_DEST_DIR ?= build

docker-build-prod:
	docker buildx build --build-arg ENVFILE=production.env -f ./app/Dockerfile -t docker-starter-prod:001 . 

docker-build:
	docker buildx build -f ./app/Dockerfile -t docker-starter:001 . 

build-app:
	go build -ldflags="-s -w" -o ./{BUILD_DEST_DIR}/docker-starter ./app

.PHONY: docker-build docker-build-prod build-app