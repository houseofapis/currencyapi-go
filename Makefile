.DEFAULT_GOAL := help
.PHONY: help
LOCAL_DOCKER_IMAGE=houseofapis/currencyapi-go
CONTAINER_NAME=currencyapi-go
WORKING_DIR=/go/app
PORT=7004
# Use official image so test/run work without building
DOCKER_IMAGE ?= golang:1.23
DOCKER_RUN = docker run --rm -v ${PWD}:${WORKING_DIR} -w ${WORKING_DIR} --name ${CONTAINER_NAME} -p ${PORT}:${PORT}
DOCKER_RUN_IT = docker run --rm -v ${PWD}:${WORKING_DIR} -w ${WORKING_DIR} --name ${CONTAINER_NAME} -p ${PORT}:${PORT} -it

build: ## Build docker image
	docker build -t ${LOCAL_DOCKER_IMAGE} . --no-cache

exec: ## Shell into the container
	${DOCKER_RUN_IT} ${DOCKER_IMAGE} sh

test: ## Run the tests (no build required)
	${DOCKER_RUN} ${DOCKER_IMAGE} go test ./...

run: ## Run the run file (no build required)
	${DOCKER_RUN_IT} ${DOCKER_IMAGE} go run example/main.go

help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
