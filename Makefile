# GOPATH:=$(shell go env GOPATH)
# VERSION=$(shell git describe --tags --always)
# INTERNAL_PROTO_FILES=$(shell find . -name *.proto)
# API_PROTO_FILES=$(shell find api -name *.proto)
# KRATOS_VERSION=$(shell go mod graph |grep go-kratos/kratos/v2 |head -n 1 |awk -F '@' '{print $$2}')
# KRATOS=$(GOPATH)/pkg/mod/github.com/go-kratos/kratos/v2@$(KRATOS_VERSION)
GO_DIR=go
BASE_FILE="info.plist icon.png"

base:
	mkdir build
	# cp $(BASE_FILE) ./build/

.PHONY: go
# build go
go:
	make base
	cd go &&	${MAKE} init
	${MAKE} -C $(GO_DIR) build
	# zip -r tools.alfredworkflow ./go/bin 

.PHONY: clean
# clean target
clean: 
	rm -rf build

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
