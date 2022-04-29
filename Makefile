## This is a self-documented Makefile. For usage information, run `make help`:
##
## For more information, refer to https://suva.sh/posts/well-documented-makefiles/

SHELL := /bin/bash

all: help

##@ Building
build: docker ##  Builds the application (same as 'docker')

set-version: ## Sets the version
	./ci/set-version.sh

docker: set-version ##  Builds the mheers/nats-seeder application
	docker buildx build --platform linux/amd64 -t mheers/nats-seeder --output type=docker .

docker-arm64: set-version ##  Builds the mheers/nats-seeder application for arm64
	docker buildx build --platform linux/arm64 -t mheers/nats-seeder --output type=docker .

docker-multi: set-version ##  Builds the mheers/nats-seeder application for amd64 and arm64
	docker buildx build --platform linux/amd64,linux/arm64 -t mheers/nats-seeder --push .

push:
	docker push mheers/nats-seeder

start:
	docker-compose up


##@ Helpers

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
