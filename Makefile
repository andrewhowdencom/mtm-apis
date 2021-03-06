# Task runner

.PHONY: help build

.DEFAULT_GOAL := help

SHELL := /bin/bash

# App version is the tag attached to this commit, if there is one. If there is not, APP_VERSION is empty. If there are
# multiple, they are concatenated using the " " (space) character.
#
# See http://stackoverflow.com/questions/4545370/how-to-list-all-tags-pointing-to-a-specific-commit-in-git
# See http://stackoverflow.com/questions/1404796/how-to-get-the-latest-tag-name-in-current-branch-in-git
APP_VERSION  := $(shell git show-ref --dereference --tags | grep ^`git rev-parse HEAD` | sed -e 's,.* refs/tags/,,' -e 's/\^{}//')
GIT_HASH     := $(shell git rev-parse --short HEAD)

#PROJECT_NS   := api
#CONTAINER_NS := api
GIT_HASH     := $(shell git rev-parse --short HEAD)

ANSI_TITLE        := '\e[1;32m'
ANSI_CMD          := '\e[0;32m'
ANSI_TITLE        := '\e[0;33m'
ANSI_SUBTITLE     := '\e[0;37m'
ANSI_WARNING      := '\e[1;31m'
ANSI_OFF          := '\e[0m'

PATH_DOCS                := $(shell pwd)/docs
PATH_BUILD_CONFIGURATION := $(shell pwd)/build

TIMESTAMP := $(shell date "+%s")

help: ## Show this menu
	@echo -e $(ANSI_TITLE)api$(ANSI_OFF)$(ANSI_SUBTITLE)" - The definitions for the meethemavdens.com APIs"$(ANSI_OFF)
	@echo -e "\nUsage: $ make \$${COMMAND} \n"
	@echo -e "Variables use the \$${VARIABLE} syntax, and are supplied as environment variables before the command. For example, \n"
	@echo -e "  \$$ VARIABLE="foo" make help\n"
	@echo -e $(ANSI_TITLE)Commands:$(ANSI_OFF)
	@grep -E '^[a-zA-Z_-%]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[32m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## Cleans up
	rm -rf \
		pkg/go/**/* \
		pkg/grpc-gateway/**/*


.PHONY: generate
generate: ## Generates the specification and libraries
	FILE='v1alpha2/services/*.proto' make generate-go
	FILE='v1alpha2/types/*.proto' make generate-go
	find . -name *.json | xargs jq --slurp '.[0] * .[1]' > api/swagger-spec/swagger.json
	echo "Done"

.PHONY: generate-go
generate-go: ## Generates the go specification
	cd api/protobuf-spec && \
	protoc -I/usr/local/include -I. \
		-I$$GOPATH/src \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:"../../pkg/go" \
		--swagger_out=logtostderr=true:../swagger-spec \
                --grpc-gateway_out="../../pkg/go" \
		$${FILE}
	mkdir -p pkg/go/$$(dirname $${FILE})
	mv pkg/go/github.com/andrewhowdencom/mtm-apis/pkg/go/$$(dirname $${FILE}) pkg/go/$$(dirname $$(dirname $${FILE}))
	rm -rf pkg/go/github.com
