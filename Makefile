.PHONY: build

build: *.go
	go build -ldflags "-s -w"

sls-deploy: build
	serverless deploy -v

fun-deploy: build
	fun deploy
