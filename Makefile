IMG ?= ghcr.io/amarchese96/mentat:latest

run:
	go run main.go

build:
	go build

docker-build:
	docker build -t ${IMG} .

docker-push:
	docker push ${IMG}