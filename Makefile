.DEFAULT_GOAL := help
.PHONY: help
LOCAL_DOCKER_IMAGE=houseofapis/currencyapi-go
CONTAINER_NAME=currencyapi-go
WORKING_DIR=/go/app
PORT=7004
DOCKER_COMMAND=docker run --rm -v ${PWD}:${WORKING_DIR} --name ${CONTAINER_NAME} -p ${PORT}:${PORT} -it ${LOCAL_DOCKER_IMAGE}

build: ## Build docker image
	docker build -t ${LOCAL_DOCKER_IMAGE} . --no-cache

exec: ## Exec into the container
	${DOCKER_COMMAND} sh

test: ## Run the tests
	${DOCKER_COMMAND} go test

run: ## Run test file
	${DOCKER_COMMAND} go run myapp/main.go

help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
