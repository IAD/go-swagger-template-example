#const
GO_SWAGGER_TAG = v0.20.1

#input
NAME ?= todo
SWAGGER_YAML_FILE ?= swagger.yml
SWAGGER_YAML_FILE_FULL := $(PWD)/${SWAGGER_YAML_FILE}
CLIENT_PKG_SUFFIX ?= client

CLIENT_PATH_FULL := $(PWD)/pkg/client
SERVER_PATH_FULL := $(PWD)/internal/server

#dynamic
MODELS_PKG_SUFFIX ?= $(CLIENT_PKG_SUFFIX)models

gen-server:
	rm -rf $(SERVER_PATH_FULL)/*
	mkdir -p $(SERVER_PATH_FULL)
	docker run --rm -it -u `id -u $(USER)` \
		-v $(PWD):$(PWD) \
		-w $(PWD) docker.io/iadolgov/go-swagger:$(GO_SWAGGER_TAG) generate server \
		-f $(SWAGGER_YAML_FILE_FULL) \
		-A $(NAME) \
		--template-dir /tmp/templates \
		--exclude-main \
		-t $(SERVER_PATH_FULL)

gen-client:
	rm -rf $(CLIENT_PATH_FULL)/*
	mkdir -p $(CLIENT_PATH_FULL)
	docker run --rm -it -u `id -u $(USER)` \
		-v $(PWD):$(PWD) \
		-w $(PWD) docker.io/iadolgov/go-swagger:$(GO_SWAGGER_TAG) generate client \
		-f $(SWAGGER_YAML_FILE_FULL) \
		-A $(NAME) \
		--template-dir /tmp/templates \
		-c $(NAME)$(CLIENT_PKG_SUFFIX) \
		-m $(NAME)$(MODELS_PKG_SUFFIX) \
		-t $(CLIENT_PATH_FULL)

lint:
	docker run --rm -v $(PWD):$(PWD) -w $(PWD) -u `id -u $(USER)` -e GOLANGCI_LINT_CACHE=/tmp/.cache -e GOCACHE=/tmp/.cache golangci/golangci-lint:v1.41.0 golangci-lint run -v --fix
